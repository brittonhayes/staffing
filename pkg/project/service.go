package project

import (
	"context"
	"errors"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
)

var _ Service = (*service)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides project methods.
type Service interface {
	CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error
	DeleteProject(ctx context.Context, command *pb.ProjectDeleteCommand) error
}

type service struct {
	projects staffing.ProjectRepository
}

func (s *service) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error {
	if command.Name == "" {
		return ErrInvalidArgument
	}

	err := s.projects.CreateProject(ctx, command.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProject(ctx context.Context, command *pb.ProjectDeleteCommand) error {
	if command.ProjectId == "" {
		return ErrInvalidArgument
	}

	err := s.projects.DeleteProject(ctx, staffing.ProjectID(command.ProjectId))
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a project service with the necessary dependencies.
func NewService(projects staffing.ProjectRepository) Service {
	return &service{
		projects: projects,
	}
}
