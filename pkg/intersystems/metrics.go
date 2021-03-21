package intersystems

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type Metric struct {
	Name    string
	Options string
	Value   float64
}

type Metrics []Metric

func (m Metrics) Frames() data.Frames {
	frame := data.NewFrame(
		"metrics",
		data.NewField("name", nil, []string{}),
		data.NewField("options", nil, []string{}),
		data.NewField("value", nil, []float64{}),
	)

	for _, v := range m {
		frame.AppendRow(
			v.Name,
			v.Options,
			v.Value,
		)
	}

	return data.Frames{frame}
}

func GetAllMetrics(ctx context.Context, client InterSystems, opts models.ListMetricsOptions) (Metrics, error) {
	var (
		metrics = []Metric{}
	)
	conn, err := client.Connect()
	if err != nil {
		return metrics, err
	}
	defer conn.Disconnect()
	var allMetrics string
	if err := conn.ClasMethod("SYS.Monitor.SAM.Sensors", "PrometheusMetrics", &allMetrics); err != nil {
		return metrics, nil
	}
	re := regexp.MustCompile(`^(\w+)({[^}]+})? (\d+|\d*\.\d+)$`)
	for _, l := range strings.Split(allMetrics, "\n") {
		vals := re.FindStringSubmatch(l)
		v, _ := strconv.ParseFloat(vals[3], 64)
		m := Metric{vals[1], vals[2], v}
		metrics = append(metrics, m)
	}

	return metrics, nil
}
