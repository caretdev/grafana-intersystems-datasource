import defaults from 'lodash/defaults';

import React, { PureComponent } from 'react';
import { InlineField, Select } from '@grafana/ui';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from './datasource';
import { defaultQuery, DataSourceOptions, QueryType, LogFile, InterSystemsQuery, LogOptions } from './types';

type Props = QueryEditorProps<DataSource, InterSystemsQuery, DataSourceOptions>;

const labelWidth = 12;
export class QueryEditor extends PureComponent<Props> {
  queryTypes: Array<SelectableValue<QueryType>> = [
    {
      label: 'Metrics',
      value: QueryType.Metrics,
      description: 'SAM Current Metrics',
    },
    {
      label: 'Log',
      value: QueryType.Log,
      description: 'Log files',
    },
  ];

  logs: Array<SelectableValue<LogFile>> = [
    {
      label: 'Alerts',
      value: LogFile.Alerts,
    },
    {
      label: 'Messages',
      value: LogFile.Messages,
    },
  ];

  onQueryTypeChange = (sel: SelectableValue<QueryType>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, queryType: sel.value! });
    onRunQuery();
  };

  onLogFileChange = (sel: SelectableValue<LogFile>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, options: { file: sel.value! } });
    onRunQuery();
  };

  render() {
    const query = defaults(this.props.query, defaultQuery);

    return (
      <>
        <InlineField label="Query type" grow={true} labelWidth={labelWidth}>
          <Select
            options={this.queryTypes}
            value={this.queryTypes.find(v => v.value === query.queryType) || this.queryTypes[0]}
            onChange={this.onQueryTypeChange}
            width={32}
          />
        </InlineField>
        {query.queryType === 'log' && (
          <InlineField label="File" labelWidth={labelWidth}>
            <Select
              options={this.logs}
              value={this.logs.find(v => v.value === (query.options as LogOptions)?.file)}
              onChange={this.onLogFileChange}
              width={32}
            />
          </InlineField>
        )}
      </>
    );
  }
}
