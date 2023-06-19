package recommend

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brittonhayes/staffing/proto/pb"
)

type loggingService struct {
	logger watermill.LoggerAdapter
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger watermill.LoggerAdapter, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateUser(ctx context.Context, command *pb.RecommendationCreateUserCommand) error {
	s.logger.Info("CreateUser", watermill.LogFields{
		"command": command,
	})

	return s.next.CreateUser(ctx, command)
}
