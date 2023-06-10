package recommend

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
		requestLatency: latency,
		next:           s,
	}
}

func (s *instrumentingService) CreateUser(ctx context.Context, command *pb.RecommendationCreateUserCommand) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateUser").Add(1)
		s.requestLatency.With("method", "CreateUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.CreateUser(ctx, command)
}
