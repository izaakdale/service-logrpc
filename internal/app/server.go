package app

import (
	"context"
	"fmt"

	logrpc "github.com/izaakdale/service-logrpc/api/v1"
	"github.com/izaakdale/service-logrpc/internal/dao"
)

type Server struct {
	client dao.LogStore
	logrpc.UnimplementedLoggingServiceServer
}

func (s *Server) Log(ctx context.Context, req *logrpc.LogRecord) (*logrpc.LogResponse, error) {
	id, err := s.client.Insert(ctx, req.Service, req.Message, req.TraceId, req.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("insert failed: %w", err)
	}
	return &logrpc.LogResponse{MessageId: id}, nil
}

func (s *Server) FetchLogs(ctx context.Context, _ *logrpc.FetchLogRequest) (*logrpc.FetchLogResponse, error) {
	recs, err := s.client.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &logrpc.FetchLogResponse{Messages: recs}, nil
}
