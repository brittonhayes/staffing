package department

import (
	"context"
	"errors"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
)

var _ Service = (*service)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides department methods.
type Service interface {
	CreateDepartment(ctx context.Context, command *pb.DepartmentCreateCommand) error
	AssignEmployee(ctx context.Context, command *pb.DepartmentAssignEmployeeCommand) error
	UnassignEmployee(ctx context.Context, command *pb.DepartmentUnassignEmployeeCommand) error
}

type service struct {
	departments staffing.DepartmentRepository
}

func (s *service) CreateDepartment(ctx context.Context, command *pb.DepartmentCreateCommand) error {
	if command.Name == "" {
		return ErrInvalidArgument
	}

	err := s.departments.CreateDepartment(ctx, command.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AssignEmployee(ctx context.Context, command *pb.DepartmentAssignEmployeeCommand) error {
	if command.DepartmentId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.departments.AssignEmployee(ctx, staffing.DepartmentID(command.DepartmentId), staffing.EmployeeID(command.EmployeeId))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UnassignEmployee(ctx context.Context, command *pb.DepartmentUnassignEmployeeCommand) error {
	if command.DepartmentId == "" || command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.departments.UnassignEmployee(ctx, staffing.DepartmentID(command.DepartmentId), staffing.EmployeeID(command.EmployeeId))
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a new project service with the necessary dependencies.
func NewService(departments staffing.DepartmentRepository) Service {
	return &service{
		departments: departments,
	}
}
