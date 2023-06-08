package department

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

func (s *loggingService) CreateDepartment(ctx context.Context, command *pb.DepartmentCreateCommand) error {
	s.logger.Info("CreateDepartment", watermill.LogFields{
		"command": command,
	})

	return s.next.CreateDepartment(ctx, command)
}

func (s *loggingService) DeleteDepartment(ctx context.Context, command *pb.DepartmentDeleteCommand) error {
	s.logger.Info("DeleteDepartment", watermill.LogFields{
		"command": command,
	})

	return s.next.DeleteDepartment(ctx, command)
}
