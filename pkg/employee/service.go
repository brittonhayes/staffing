package employee

import (
	"context"
	"errors"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
)

var _ Service = (*service)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides employee methods.
type Service interface {
	CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) (*staffing.Employee, error)
	DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error
	AssignProject(ctx context.Context, command *pb.EmployeeAssignProjectCommand) error
	UnassignProject(ctx context.Context, command *pb.EmployeeUnassignProjectCommand) error
}

type service struct {
	employees staffing.EmployeeRepository
}

func (s *service) CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) (*staffing.Employee, error) {

	if err := command.ValidateAll(); err != nil {
		return nil, err
	}

	resp, err := s.employees.CreateEmployee(ctx, command.Name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error {

	err := s.employees.DeleteEmployee(ctx, staffing.EmployeeID(command.Id))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AssignProject(ctx context.Context, command *pb.EmployeeAssignProjectCommand) error {

	err := s.employees.AssignProject(ctx, staffing.ProjectID(command.ProjectId), staffing.EmployeeID(command.Id))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UnassignProject(ctx context.Context, command *pb.EmployeeUnassignProjectCommand) error {

	err := s.employees.UnassignProject(ctx, staffing.EmployeeID(command.Id))
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a new project service with the necessary dependencies.
func NewService(employees staffing.EmployeeRepository) Service {
	return &service{
		employees: employees,
	}
}
