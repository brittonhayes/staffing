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

func (s *instrumentingService) DeleteDepartment(ctx context.Context, command *pb.DepartmentDeleteCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteDepartment").Add(1)
		s.requestLatency.With("method", "DeleteDepartment").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.DeleteDepartment(ctx, command)
}
