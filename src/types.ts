import { DataQuery, DataSourceJsonData } from '@grafana/data';

export enum QueryType {
  Metrics = 'metrics',
  Alerts = 'alerts',
}

export interface InterSystemsQuery extends Metrics {

}

export interface Metrics extends DataQuery {
}

export const defaultQuery: Partial<Metrics> = {
  queryType: QueryType.Metrics,
};

/**
 * These are options configured for each DataSource instance
 */
export interface DataSourceOptions extends DataSourceJsonData {}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface SecureJsonData {
  password?: string;
}
