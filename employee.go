package staffing

import (
	"context"
	"errors"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, name string) (*Employee, error)
	DeleteEmployee(ctx context.Context, employeeID EmployeeID) error

	AssignProject(ctx context.Context, projectID ProjectID, employeeID EmployeeID) error
	UnassignProject(ctx context.Context, employeeID EmployeeID) error

	AssignDepartment(ctx context.Context, departmentID DepartmentID, employeeID EmployeeID) error
	UnassignDepartment(ctx context.Context, employeeID EmployeeID) error

	Close() error
}

type EmployeeID string

type Employee struct {
	ID                 EmployeeID   `json:"id"`
	Name               string       `json:"name"`
	AssignedProject    ProjectID    `json:"assigned_project"`
	AssignedDepartment DepartmentID `json:"assigned_department"`
}
