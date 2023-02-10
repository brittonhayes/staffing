package inmem

import (
	"context"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/staffing"
)

type departmentRepository struct {
	mu          sync.RWMutex
	departments map[staffing.DepartmentID]*staffing.Department
}

func NewDepartmentRepository() staffing.DepartmentRepository {
	return &departmentRepository{
		departments: make(map[staffing.DepartmentID]*staffing.Department),
	}
}

func (r *departmentRepository) Close() error {
	defer r.mu.Unlock()
	return nil
}

func (r *departmentRepository) CreateDepartment(ctx context.Context, name string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.departments[staffing.DepartmentID(name)] = &staffing.Department{
		ID:   staffing.DepartmentID(name),
		Name: name,
	}

	return nil
}

func (r *departmentRepository) AssignEmployee(ctx context.Context, departmentID staffing.DepartmentID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.departments[departmentID].Employees = append(r.departments[departmentID].Employees, staffing.Employee{
		ID:   employeeID,
		Name: gofakeit.Name(),
	})

	return nil
}

func (r *departmentRepository) UnassignEmployee(ctx context.Context, departmentID staffing.DepartmentID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}
