import { DataSourceInstanceSettings } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { DataSourceOptions, InterSystemsQuery } from './types';

export class DataSource extends DataSourceWithBackend<InterSystemsQuery, DataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceOptions>) {
    super(instanceSettings);
  }
}
