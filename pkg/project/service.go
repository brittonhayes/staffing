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
	AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error
	UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error
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

func (s *service) AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error {
	if command.ProjectId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.projects.AssignEmployee(ctx, staffing.ProjectID(command.ProjectId), staffing.EmployeeID(command.EmployeeId))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error {
	if command.ProjectId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.projects.UnassignEmployee(ctx, staffing.ProjectID(command.ProjectId), staffing.EmployeeID(command.EmployeeId))
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
