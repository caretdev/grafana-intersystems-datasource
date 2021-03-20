import React, { PureComponent } from 'react';
import { LegacyForms, Input } from '@grafana/ui';
import { onUpdateDatasourceOption, DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { DataSourceOptions, SecureJsonData } from './types';

const { SecretFormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<DataSourceOptions> {}

interface State {}

export class ConfigEditor extends PureComponent<Props, State> {
  onSettingReset = (prop: string) => (event: any) => {
    this.onSettingUpdate(prop, false)({ target: { value: undefined } });
  };

  onSettingUpdate = (prop: string, set = true) => (event: any) => {
    const { onOptionsChange, options } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: {
        ...options.secureJsonData,
        [prop]: event.target.value,
      },
      secureJsonFields: {
        ...options.secureJsonFields,
        [prop]: set,
      },
    });
  };

  render() {
    const {
      options: { secureJsonData, secureJsonFields, url, user, database },
    } = this.props;
    const secureSettings = (secureJsonData || {}) as SecureJsonData;

    return (
      <>
        <h3 className="page-heading">InterSystems IRISConnection</h3>

        <div className="gf-form-group">
          <div className="gf-form max-width-30">
            <span className="gf-form-label width-10">Host</span>
            <Input
              css=""
              value={url}
              placeholder="localhost:1972"
              onChange={onUpdateDatasourceOption(this.props, 'user')}
              required
            />
          </div>

          <div className="gf-form max-width-30">
            <span className="gf-form-label width-10">Namespace</span>
            <Input
              css=""
              value={database}
              placeholder="USER"
              onChange={onUpdateDatasourceOption(this.props, 'database')}
            />
          </div>

          <div className="gf-form-inline">
            <div className="gf-form max-width-20">
              <span className="gf-form-label width-10">Login</span>
              <Input
                css=""
                value={user}
                placeholder=""
                onChange={onUpdateDatasourceOption(this.props, 'user')}
                required
              />
            </div>
            <div className="gf-form">
              <SecretFormField
                isConfigured={secureJsonFields!['password']}
                value={secureSettings.password}
                onReset={this.onSettingReset('password')}
                onChange={this.onSettingUpdate('password', false)}
                onBlur={this.onSettingUpdate('password')}
                inputWidth={9}
              />
            </div>
          </div>
        </div>
      </>
    );
  }
}
