package staffing

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, name string) (*Employee, error)
	DeleteEmployee(ctx context.Context, employeeID EmployeeID) error

	AssignProject(ctx context.Context, employeeID EmployeeID, projectID ProjectID) error
	UnassignProject(ctx context.Context, employeeID EmployeeID) error

	AssignDepartment(ctx context.Context, employeeID EmployeeID, departmentID DepartmentID) error
	UnassignDepartment(ctx context.Context, employeeID EmployeeID) error

	InsertFeedback(ctx context.Context, feedback *Feedback) error

	Close() error
}

type EmployeeID string

type Employee struct {
	bun.BaseModel `bun:"table:employees"`

	ID                 EmployeeID   `json:"id" bun:",pk"`
	Name               string       `json:"name"`
	AssignedProject    ProjectID    `json:"assigned_project"`
	AssignedDepartment DepartmentID `json:"assigned_department"`
}

type Feedback struct {
	EmployeeID   EmployeeID `json:"employee_id"`
	FeedbackType string     `json:"feedback_type"`
	ItemID       string     `json:"item_id"`
	Timestamp    time.Time  `json:"timestamp"`
}

func NewEmployee(name string) *Employee {
	return &Employee{
		ID:                 EmployeeID(uuid.NewString()),
		Name:               name,
		AssignedProject:    "",
		AssignedDepartment: "",
	}
}

func NewFeedback(employeeID EmployeeID, feedbackType string, itemID string, timestamp time.Time) *Feedback {
	return &Feedback{
		EmployeeID:   employeeID,
		FeedbackType: feedbackType,
		ItemID:       itemID,
		Timestamp:    timestamp,
	}
}
