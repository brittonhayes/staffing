package project

import (
	"context"
	"time"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) (*staffing.Project, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateProject").Add(1)
		s.requestLatency.With("method", "CreateProject").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CreateProject(ctx, command)
}

func (s *instrumentingService) CancelProject(ctx context.Context, command *pb.ProjectCancelCommand) (*staffing.Project, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteProject").Add(1)
		s.requestLatency.With("method", "DeleteProject").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CancelProject(ctx, command)
}
