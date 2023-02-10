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

func (s *loggingService) AssignEmployee(ctx context.Context, command *pb.DepartmentAssignEmployeeCommand) error {
	s.logger.Info("AssignEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.AssignEmployee(ctx, command)
}

func (s *loggingService) UnassignEmployee(ctx context.Context, command *pb.DepartmentUnassignEmployeeCommand) error {
	s.logger.Info("UnassignEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.UnassignEmployee(ctx, command)
}
