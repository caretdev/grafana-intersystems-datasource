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
		Name: query.Options.Name,
	}
	if req.RefID == "listMetrics" {
		return GetListMetrics(ctx, d.client, opt)
	}
 	return GetAllMetrics(ctx, d.client, opt)
}

func (d *Datasource) HandleLogQuery(ctx context.Context, query *models.LogQuery, req backend.DataQuery) (dfutil.Framer, error) {
	opt := models.ListLogOptions {
		File: query.Options.File,
	}
	return GetLog(ctx, d.client, opt)
}

func (d *Datasource) HandleApplicationErrorsQuery(ctx context.Context, query *models.ApplicationErrorsQuery, req backend.DataQuery) (dfutil.Framer, error) {
	opt := models.ListApplicationErrorsOptions {
	}
	return GetApplicationErrors(ctx, d.client, opt)
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
	addr, namespace, user, password := settings.Addr, settings.Namespace, settings.User, settings.Password
	return &Datasource{NewInterSystems(addr, namespace, user, password)}
}
