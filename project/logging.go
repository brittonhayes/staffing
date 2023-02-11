package project

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brittonhayes/staffing/proto/pb"
)

type loggingService struct {
	logger watermill.LoggerAdapter
	next   CommandService
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger watermill.LoggerAdapter, s CommandService) CommandService {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error {
	s.logger.Info("CreateProject", watermill.LogFields{
		"command": command,
	})

	return s.next.CreateProject(ctx, command)
}

func (s *loggingService) AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error {
	s.logger.Info("AssignEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.AssignEmployee(ctx, command)
}

func (s *loggingService) UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error {
	s.logger.Info("UnassignEmployee", watermill.LogFields{
		"command": command,
	})

	return s.next.UnassignEmployee(ctx, command)
}
