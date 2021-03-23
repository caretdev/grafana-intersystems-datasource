package plugin

import (
	"context"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/dfutil"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// The Datasource type handles the requests sent to the datasource backend
type Datasource interface {
	HandleMetricsQuery(context.Context, *models.MetricsQuery, backend.DataQuery) (dfutil.Framer, error)
	HandleLogQuery(context.Context, *models.LogQuery, backend.DataQuery) (dfutil.Framer, error)
	CheckHealth(context.Context) error
}

// HandleQueryData handles the `QueryData` request for the datasource
func HandleQueryData(ctx context.Context, d Datasource, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	m := GetQueryHandlers(&Server{
		Datasource: d,
	})

	return m.QueryData(ctx, req)
}

// CheckHealth ensures that the datasource settings are able to retrieve data from 
func CheckHealth(ctx context.Context, d Datasource, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	if err := d.CheckHealth(ctx); err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}
	return &backend.CheckHealthResult{
		Status: backend.HealthStatusOk,
		Message: backend.HealthStatusOk.String(),
	}, nil
}
