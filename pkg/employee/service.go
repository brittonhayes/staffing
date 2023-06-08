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
	CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) error
	DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error
}

type service struct {
	employees staffing.EmployeeRepository
}

func (s *service) CreateEmployee(ctx context.Context, command *pb.EmployeeCreateCommand) error {
	if command.Name == "" {
		return ErrInvalidArgument
	}

	err := s.employees.CreateEmployee(ctx, command.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteEmployee(ctx context.Context, command *pb.EmployeeDeleteCommand) error {
	if command.EmployeeId == "" {
		return ErrInvalidArgument
	}

	err := s.employees.DeleteEmployee(ctx, staffing.EmployeeID(command.EmployeeId))
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
