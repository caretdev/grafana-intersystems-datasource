package plugin

import (
	"context"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/intersystems"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

// Instance is the root Datasource implementation that wraps a Datasource
type Instance struct {
	Datasource Datasource
}

// CheckHealth ...
func (i *Instance) CheckHealth(ctx context.Context) error {
	return i.Datasource.CheckHealth(ctx)
}

// NewInstance creates a new Instance using the settings to determine if things like the Caching Wrapper should be enabled
func NewInstance(ctx context.Context, settings models.Settings) *Instance {
	var (
		isc = intersystems.NewDatasource(ctx, settings)
	)

	var d Datasource = isc

	// TODO: wrap these HTTP handlers with a caching wrapper
	return &Instance{
		Datasource: d,
	}
}

func newDataSourceInstance(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	datasourceSettings, err := models.LoadSettings(settings)
	if err != nil {
		return nil, err
	}

	return NewInstance(context.Background(), datasourceSettings), nil
}
