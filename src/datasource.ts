import {
  DataSourceInstanceSettings,
  DataQueryRequest,
  DataFrameView,
  DataFrame,
  DataQueryResponse,
} from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { DataSourceOptions, InterSystemsQuery } from './types';

export class DataSource extends DataSourceWithBackend<InterSystemsQuery, DataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceOptions>) {
    super(instanceSettings);
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
