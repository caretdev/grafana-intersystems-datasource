import defaults from 'lodash/defaults';

import React, { PureComponent } from 'react';
import { InlineField, Select } from '@grafana/ui';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from './datasource';
import { defaultQuery, DataSourceOptions, QueryType, Metrics } from './types';

type Props = QueryEditorProps<DataSource, Metrics, DataSourceOptions>;

const labelWidth = 12;
export class QueryEditor extends PureComponent<Props> {
  queryTypes: Array<SelectableValue<QueryType>> = [
    {
      label: 'Metrics',
      value: QueryType.Metrics,
      description: 'SAM Current Metrics',
    },
    {
      label: 'Alerts',
      value: QueryType.Alerts,
      description: 'SAM Latest Alerts',
    },
  ];

  onQueryTypeChange = (sel: SelectableValue<QueryType>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, queryType: sel.value! });
    onRunQuery();
  };

  render() {
    const query = defaults(this.props.query, defaultQuery);

    return (
      <div className="gf-form">
        <InlineField label="Query type" grow={true} labelWidth={labelWidth}>
          <Select
            options={this.queryTypes}
            value={this.queryTypes.find(v => v.value === query.queryType) || this.queryTypes[0]}
            onChange={this.onQueryTypeChange}
          />
        </InlineField>
      </div>
    );
  }
}
