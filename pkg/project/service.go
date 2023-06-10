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
	CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) (*staffing.Project, error)
	CancelProject(ctx context.Context, command *pb.ProjectCancelCommand) (*staffing.Project, error)
}

type service struct {
	projects staffing.ProjectRepository
}

func (s *service) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) (*staffing.Project, error) {

	if err := command.ValidateAll(); err != nil {
		return nil, err
	}

	resp, err := s.projects.CreateProject(ctx, command.Name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) CancelProject(ctx context.Context, command *pb.ProjectCancelCommand) (*staffing.Project, error) {

	if err := command.ValidateAll(); err != nil {
		return nil, err
	}

	resp, err := s.projects.CancelProject(ctx, staffing.ProjectID(command.Id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NewService creates a project service with the necessary dependencies.
func NewService(projects staffing.ProjectRepository) Service {
	return &service{
		projects: projects,
	}
}
