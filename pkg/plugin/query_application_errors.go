package plugin

import (
	"context"

	"github.com/caretdev/grafana-intersystems-datasource/pkg/dfutil"
	"github.com/caretdev/grafana-intersystems-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func (s *Server) handleApplicationErrorsQuery(ctx context.Context, q backend.DataQuery) backend.DataResponse {
	query := &models.ApplicationErrorsQuery{}
	if err := UnmarshalQuery(q.JSON, query); err != nil {
		return *err
	}
	return dfutil.FrameResponseWithError(s.Datasource.HandleApplicationErrorsQuery(ctx, query, q))
}

func (s *Server) HandleApplicationErrors(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	return &backend.QueryDataResponse{
		Responses: processQueries(ctx, req, s.handleApplicationErrorsQuery),
	}, nil
}
