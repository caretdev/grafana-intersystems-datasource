package intersystems

import (
	"context"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/dfutil"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type Datasource struct {
	client InterSystems
}

func (d *Datasource) HandleMetricsQuery(ctx context.Context, query *models.MetricsQuery, req backend.DataQuery) (dfutil.Framer, error) {
	opt := models.ListMetricsOptions{
	}
	return GetAllMetrics(ctx, d.client, opt)
}

func (d *Datasource) HandleAlertsQuery(ctx context.Context, query *models.AlertsQuery, req backend.DataQuery) (dfutil.Framer, error) {
	return nil, nil
}

func (d *Datasource) CheckHealth(ctx context.Context) error {
	var conn, err = d.client.Connect()
	if err != nil {
		return err
	}
	defer conn.Disconnect()
	return err
}

func NewDatasource(ctx context.Context, settings models.Settings) *Datasource {
	addr, user, password := settings.Addr, settings.User, settings.Password
	namespace := "%SYS"
	return &Datasource{NewInterSystems(addr, namespace, user, password)}
}
