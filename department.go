package staffing

import (
	"context"
)

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, name string) error
	DeleteDepartment(ctx context.Context, departmentID DepartmentID) error
	Close() error
}

type DepartmentID string

type Department struct {
	ID   DepartmentID `json:"id"`
	Name string       `json:"name"`
}
