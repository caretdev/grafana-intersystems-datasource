package intersystems

import (
	"context"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/caretdev/go-irisnative/src/connection"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var re = regexp.MustCompile(`^(\w+)(?:{(?:id="([^"]+)")?[^}]*})? (\d+|\d*\.\d+)$`)

type Metric struct {
	Name  string
	Id    string
	Value float64
}

type Metrics struct {
	metrics []Metric
	split   bool
	ts      time.Time
}

func (ms Metrics) Frames() data.Frames {
	frames := []*data.Frame{}

	if ms.split {
		for _, m := range ms.metrics {

			frame := data.NewFrame("",
				data.NewField("Time", nil, []time.Time{ms.ts}),
				data.NewField(m.Id, nil, []float64{m.Value}),
			).SetMeta(&data.FrameMeta{
				PreferredVisualization: "table",
			})
			frames = append(frames, frame)
		}
	} else {
		frame := data.NewFrame("",
			data.NewField("Time", nil, []time.Time{}),
			data.NewField("Name", nil, []string{}),
			data.NewField("Id", nil, []string{}),
			data.NewField("Value", nil, []float64{}),
		).SetMeta(&data.FrameMeta{
			PreferredVisualization: "table",
		})
		frames = append(frames, frame)

		for _, m := range ms.metrics {
			frame.AppendRow(
				ms.ts,
				m.Name,
				m.Id,
				m.Value,
			)
		}
	}

	return frames
}

type ListMetrics map[string]bool

func (l ListMetrics) Frames() data.Frames {
	metrics := make([]string, 0)
	for k := range l {
		metrics = append(metrics, k)
	}
	sort.Strings(metrics)
	frame := data.NewFrame("",
		data.NewField("name", nil, metrics),
	)

	return data.Frames{frame}
}

func GetListMetrics(ctx context.Context, client InterSystems, opts models.ListMetricsOptions) (ListMetrics, error) {
	var list = make(map[string]bool)
	var ml Metrics
	var err error
	if ml, err = GetAllMetrics(ctx, client, opts); err != nil {
		return list, err
	}
	for _, metric := range ml.metrics {
		list[metric.Name] = true
	}
	return list, nil
}

func GetAllMetrics(ctx context.Context, client InterSystems, opts models.ListMetricsOptions) (Metrics, error) {
	var (
		metrics = []Metric{}
		split   = opts.Name != ""
		ts      = time.Now()
	)
	var conn connection.Connection
	var err error
	if conn, err = client.ConnectSYS(); err != nil {
		return Metrics{metrics, split, ts}, err
	}
	defer conn.Disconnect()
	var allMetrics string
	if err := conn.ClassMethod("SYS.Monitor.SAM.Sensors", "PrometheusMetrics", &allMetrics); err != nil {
		return Metrics{metrics, split, ts}, nil
	}
	for _, l := range strings.Split(allMetrics, "\n") {
		vals := re.FindStringSubmatch(l)
		name, id, value := vals[1], vals[2], vals[3]
		if opts.Name != "" && name != opts.Name {
			continue
		}
		v, _ := strconv.ParseFloat(value, 64)
		if id == "" {
			id = "Value"
		}
		m := Metric{name, id, v}
		metrics = append(metrics, m)
	}

	return Metrics{metrics, split, ts}, nil
}
