package employee

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

func (s *loggingService) CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) error {
	s.logger.Info("CreateEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.CreateEmployee(ctx, command)
}

func (s *loggingService) DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error {
	s.logger.Info("DeleteEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.DeleteEmployee(ctx, command)
}
