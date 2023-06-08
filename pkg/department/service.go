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
	DeleteDepartment(ctx context.Context, command *pb.DepartmentDeleteCommand) error
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

func (s *service) DeleteDepartment(ctx context.Context, command *pb.DepartmentDeleteCommand) error {
	if command.DepartmentId == "" {
		return ErrInvalidArgument
	}

	err := s.departments.DeleteDepartment(ctx, staffing.DepartmentID(command.DepartmentId))
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
