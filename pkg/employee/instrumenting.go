package employee

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

func (s *instrumentingService) CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateEmployee").Add(1)
		s.requestLatency.With("method", "CreateEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CreateEmployee(ctx, command)
}

func (s *instrumentingService) DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteEmployee").Add(1)
		s.requestLatency.With("method", "DeleteEmployee").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.DeleteEmployee(ctx, command)
}
