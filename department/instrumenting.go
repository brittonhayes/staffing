package department

import (
	"context"
	"time"

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

func (s *instrumentingService) CreateDepartment(ctx context.Context, command *pb.DepartmentCreateCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateDepartment").Add(1)
		s.requestLatency.With("method", "CreateDepartment").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CreateDepartment(ctx, command)
}

func (s *instrumentingService) AssignEmployee(ctx context.Context, command *pb.DepartmentAssignEmployeeCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AssignEmployee").Add(1)
		s.requestLatency.With("method", "AssignEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.AssignEmployee(ctx, command)
}

func (s *instrumentingService) UnassignEmployee(ctx context.Context, command *pb.DepartmentUnassignEmployeeCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UnassignEmployee").Add(1)
		s.requestLatency.With("method", "UnassignEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.UnassignEmployee(ctx, command)
}
