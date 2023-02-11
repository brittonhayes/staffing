package project

import (
	"context"
	"errors"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
)

var _ CommandService = (*commandService)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

// CommandService is the interface that provides write-only project methods.
type CommandService interface {
	CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error
	AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error
	UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error
}

type commandService struct {
	projects staffing.ProjectRepository
}

func (s *commandService) CreateProject(ctx context.Context, command *pb.ProjectCreateCommand) error {
	if command.Name == "" {
		return ErrInvalidArgument
	}

	err := s.projects.CreateProject(ctx, command.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *commandService) AssignEmployee(ctx context.Context, command *pb.ProjectAssignEmployeeCommand) error {
	if command.ProjectId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.projects.AssignEmployee(ctx, staffing.ProjectID(command.ProjectId), staffing.EmployeeID(command.EmployeeId))
	if err != nil {
		return err
	}

	return nil
}

func (s *commandService) UnassignEmployee(ctx context.Context, command *pb.ProjectUnassignEmployeeCommand) error {
	if command.ProjectId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.projects.UnassignEmployee(ctx, staffing.ProjectID(command.ProjectId), staffing.EmployeeID(command.EmployeeId))
	if err != nil {
		return err
	}

	return nil
}

// NewCommandService creates a project service with the necessary dependencies.
func NewCommandService(projects staffing.ProjectRepository) CommandService {
	return &commandService{
		projects: projects,
	}
}
