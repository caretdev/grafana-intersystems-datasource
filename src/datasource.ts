import { Observable, merge } from 'rxjs';
import {
  DataSourceInstanceSettings,
  DataQueryRequest,
  DataFrameView,
  DataFrame,
  LoadingState,
  CircularDataFrame,
  DataQueryResponse,
} from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { DataSourceOptions, InterSystemsQuery, MetricsOptions } from './types';

export class DataSource extends DataSourceWithBackend<InterSystemsQuery, DataSourceOptions> {
  private frames: Map<string, CircularDataFrame> = new Map<string, CircularDataFrame>();

  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceOptions>) {
    super(instanceSettings);
  }

  query(options: DataQueryRequest<InterSystemsQuery>): Observable<DataQueryResponse> {
    const streams = options.targets.map((target) => {
      if (target.queryType === 'metrics' && target.refId !== 'listMetrics') {
        return this.streamMetrics(target, options);
      }
      return super.query({ ...options, targets: [target] });
    });

    return merge(...streams);
  }

  streamMetrics(
    target: InterSystemsQuery,
    options: DataQueryRequest<InterSystemsQuery>
  ): Observable<DataQueryResponse> {
    const queryOptions = target.options as MetricsOptions;
    return new Observable<DataQueryResponse>((subscriber) => {
      const intervalId = setInterval(() => {
        super.query({ ...options, targets: [target] }).forEach((value) => {
          const frames: CircularDataFrame[] = value.data.map((dataFrame: DataFrame, ind) => {
            const frameKey = `${queryOptions.name}-${ind}`;
            const frame = this.frames.get(frameKey);
            if (!frame) {
              const frame = new CircularDataFrame({
                append: 'tail',
                capacity: 1000,
              });
              dataFrame.fields.forEach((field) => frame.addField(field));
              this.frames.set(frameKey, frame);
              return frame;
            }
            const value: { [key: string]: Object[] } = {};
            dataFrame.fields.forEach((field) => {
              value[field.name || 'value'] = field.values.get(0);
            });
            frame.add(value);
            return frame;
          });
          subscriber.next({
            data: frames,
            key: target.refId,
            state: LoadingState.Streaming,
          });
        });
      }, options.intervalMs);

      return () => {
        clearInterval(intervalId);
      };
    });
  }

  async getListMetrics(): Promise<string[]> {
    const request = {
      targets: [
        {
          queryType: 'metrics',
          refId: 'listMetrics',
        },
      ],
    } as DataQueryRequest;

    let res: DataQueryResponse;

    try {
      res = await this.query(request).toPromise();
    } catch (err) {
      return Promise.reject(err);
    }

    if (!res || !res.data || res.data.length < 0) {
      return [];
    }

    const view = new DataFrameView(res.data[0] as DataFrame);
    return view.map((item) => {
      return item['name'];
    });
  }
}
