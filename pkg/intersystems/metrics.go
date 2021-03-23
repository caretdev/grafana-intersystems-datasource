package intersystems

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/caretdev/go-irisnative/src/connection"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type Metric struct {
	Name  string
	Id    string
	Value float64
}

type Metrics []Metric

func (ms Metrics) Frames() data.Frames {
	frame := data.NewFrame("",
		data.NewField("name", nil, []string{}),
		data.NewField("id", nil, []string{}),
		data.NewField("value", nil, []float64{}),
	)

	for _, m := range ms {
		frame.AppendRow(
			m.Name, 
			m.Id,
			m.Value,
		)
	} 

	return data.Frames{frame}
}

func GetAllMetrics(ctx context.Context, client InterSystems, opts models.ListMetricsOptions) (Metrics, error) {
	var (
		metrics = []Metric{}
	)
	var conn connection.Connection
	var err error
	if conn, err = client.Connect(); err != nil {
		return metrics, err
	}
	defer conn.Disconnect()
	var allMetrics string
	if err := conn.ClassMethod("SYS.Monitor.SAM.Sensors", "PrometheusMetrics", &allMetrics); err != nil {
		return metrics, nil
	}
	re := regexp.MustCompile(`^(\w+)(?:{(?:id="([^"]+)")?[^}]*})? (\d+|\d*\.\d+)$`)
	for _, l := range strings.Split(allMetrics, "\n") {
		// if !strings.Contains(l, "iris_cpu_pct") {
		// 	continue
		// }
		vals := re.FindStringSubmatch(l)
		name, id, value := vals[1], vals[2], vals[3]
		v, _ := strconv.ParseFloat(value, 64)
		m := Metric{name, id, v}
		metrics = append(metrics, m)
	}

	return metrics, nil
}
