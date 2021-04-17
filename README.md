# InterSystems IRIS Data Source

[![Build](https://github.com/caretdev/grafana-intersystems-datasource/workflows/CI/badge.svg)](https://github.com/caretdev/grafana-intersystems-datasource/actions?query=workflow%3A%22CI%22)

This is Grafana data source for showing metrics from InterSystems IRIS

## Features

- Utilizes InterSystems xDBC protocol
- Streaming of SAM Metrics in realtime updated with interval set in Grafana, with history
- Shows Logs: Alerts and Messages
- Application Errors (stored in ^ERRORS)

## Testing

- Clone this repo
- `docker-compose up -d`
- open localhost:3000

## Configuration

With docker-compose it will be provisioned automatically, and will only allow to test the connection.

![NewDataSource](https://raw.githubusercontent.com/caretdev/grafana-intersystems-datasource/main/img/configuration.png)

## Explore

![Metrics](https://raw.githubusercontent.com/caretdev/grafana-intersystems-datasource/main/img/metrics.png)

## Panels

![MetricsPanel](https://raw.githubusercontent.com/caretdev/grafana-intersystems-datasource/main/img/MetricsPanel.png)
![LogsPanel](https://raw.githubusercontent.com/caretdev/grafana-intersystems-datasource/main/img/LogsPanel.png)
![ErrorsPanel](https://raw.githubusercontent.com/caretdev/grafana-intersystems-datasource/main/img/ErrorsPanel.png)

## Learn more

- [InterSystems](https://intersystems.com)
- [InterSystems IRIS](https://www.intersystems.com/products/intersystems-iris/)
- [InterSystems Developer Community](https://community.intersystems.com/)
- [InterSystems Support](https://www.intersystems.com/support-learning/support/)
- [InterSystems Learning](https://learning.intersystems.com/)
