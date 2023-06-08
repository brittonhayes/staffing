package inmem

import (
	"context"
	"sync"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
)

type employeeRepository struct {
	mu        sync.RWMutex
	employees map[staffing.EmployeeID]*staffing.Employee
}

func NewEmployeeRepository() staffing.EmployeeRepository {
	return &employeeRepository{
		employees: make(map[staffing.EmployeeID]*staffing.Employee),
	}
}

func (r *employeeRepository) Close() error {
	defer r.mu.Unlock()
	return nil
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, name string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id := staffing.EmployeeID(uuid.NewString())
	r.employees[id] = &staffing.Employee{
		ID:   id,
		Name: name,
	}

	return nil
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employeeID staffing.EmployeeID) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id := staffing.EmployeeID(uuid.NewString())
	r.employees[id] = nil

	return nil
}

func (r *employeeRepository) AssignProject(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.employees[employeeID].AssignedProject = projectID

	return nil
}

func (r *employeeRepository) UnassignProject(ctx context.Context, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.employees[employeeID].AssignedProject = ""

	return nil
}

func (r *employeeRepository) AssignDepartment(ctx context.Context, departmentID staffing.DepartmentID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.employees[employeeID].AssignedDepartment = departmentID

	return nil
}

func (r *employeeRepository) UnassignDepartment(ctx context.Context, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.employees[employeeID].AssignedDepartment = ""

	return nil
}