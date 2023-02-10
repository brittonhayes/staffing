package staffing

import (
	"context"
)

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, name string) error
	AssignEmployee(ctx context.Context, departmentID DepartmentID, employeeID EmployeeID) error
	UnassignEmployee(ctx context.Context, departmentID DepartmentID, employeeID EmployeeID) error
	Close() error
}

type DepartmentID string

type Department struct {
	ID        DepartmentID `json:"id"`
	Name      string       `json:"name"`
	Employees []Employee   `json:"employees"`
}
