package staffing

import (
	"context"
	"errors"
)

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	CreateProject(ctx context.Context, name string) error
	AssignEmployee(ctx context.Context, projectID ProjectID, employeeID EmployeeID) error
	UnassignEmployee(ctx context.Context, projectID ProjectID, employeeID EmployeeID) error
	Close() error
}

type ProjectID string

type Project struct {
	ID                ProjectID
	Name              string
	AssignedEmployees []EmployeeID
}

type EmployeeID string

type Employee struct {
	ID   EmployeeID
	Name string
}
