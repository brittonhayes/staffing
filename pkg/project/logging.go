package project

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brittonhayes/staffing"
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

func (s *loggingService) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) (*staffing.Project, error) {
	s.logger.Info("CreateProject", watermill.LogFields{
		"command": command,
	})

	return s.next.CreateProject(ctx, command)
}

func (s *loggingService) CancelProject(ctx context.Context, command *pb.ProjectCancelCommand) (*staffing.Project, error) {
	s.logger.Info("CancelProject", watermill.LogFields{
		"command": command,
	})

	return s.next.CancelProject(ctx, command)
}
