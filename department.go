package staffing

import "context"

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, name string) error
	AssignEmployee(ctx context.Context, departmentID DepartmentID, employeeID EmployeeID) error
	UnassignEmployee(ctx context.Context, departmentID DepartmentID, employeeID EmployeeID) error
}

type DepartmentID string

type Department struct {
	ID   DepartmentID
	Name string
}
