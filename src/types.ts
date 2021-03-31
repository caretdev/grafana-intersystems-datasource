import { DataQuery, DataSourceJsonData } from '@grafana/data';

export enum QueryType {
  Metrics = 'metrics',
  Log = 'log',
  ApplicationErrors = 'application_errors',
}

export interface Metrics extends DataQuery {}

export const defaultQuery: Partial<Metrics> = {
  queryType: QueryType.Metrics,
};

export interface MetricsOptions {
  name?: string;
}

export interface LogOptions {
  file?: string;
}
export interface InterSystemsQuery extends DataQuery {
  options?: MetricsOptions | LogOptions;
}

export enum LogFile {
  Messages = 'messages.log',
  Alerts = 'alerts.log',
}

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
