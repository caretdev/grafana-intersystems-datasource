package intersystems

import (
	"context"
	"fmt"

	"github.com/caretdev/go-irisnative/src/connection"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type Datasource struct {
	err error
}

func (d *Datasource) CheckHealth(ctx context.Context) error {
	return d.err
}

func NewDatasource(ctx context.Context, settings models.Settings) *Datasource {
	namespace := settings.Namespace
	if len(namespace) == 0 {
		namespace = "%SYS"
	}
	log.DefaultLogger.Debug(fmt.Sprintf("Check connection to InterSystems: %v@%v/%v/",settings.User, settings.Addr, settings.Namespace))
	var _, err = connection.Connect(settings.Addr, namespace, settings.User, settings.Password)
	// defer conn.Disconnect()
	log.DefaultLogger.Debug("Check connection to InterSystems: OK")
	return &Datasource{err}
}
