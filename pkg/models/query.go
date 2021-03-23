package models

const (
	QueryTypeMetrics = "metrics"
	QueryTypeLog = "log"
)

type MetricsQuery struct {
	Options ListMetricsOptions `json:"options"`
}

type LogQuery struct {
	Options ListLogOptions `json:"options"`
}