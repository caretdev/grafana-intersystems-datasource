package models

const (
	QueryTypeMetrics = "metrics"
	QueryTypeLog = "log"
	QueryTypeApplicationErrors = "application_errors"
)

type ListMetricsOptions struct {
	Name  string `json:"name"`
}

type MetricsQuery struct {
	Options ListMetricsOptions `json:"options"`
}

type ListLogOptions struct {
	File string `json:"file"`
}

type LogQuery struct {
	Options ListLogOptions `json:"options"`
}

type ListApplicationErrorsOptions struct {
}

type ApplicationErrorsQuery struct {
	Options ListApplicationErrorsOptions `json:"options"`
}