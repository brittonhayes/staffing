package project

import (
	"context"
	"time"

	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           CommandService
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s CommandService) CommandService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateProject").Add(1)
		s.requestLatency.With("method", "CreateProject").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CreateProject(ctx, command)
}

func (s *instrumentingService) AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AssignEmployee").Add(1)
		s.requestLatency.With("method", "AssignEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.AssignEmployee(ctx, command)
}

func (s *instrumentingService) UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UnassignEmployee").Add(1)
		s.requestLatency.With("method", "UnassignEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.UnassignEmployee(ctx, command)
}
