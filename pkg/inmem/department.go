package inmem

import (
	"context"
	"sync"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
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

	id := staffing.DepartmentID(uuid.NewString())
	r.departments[id] = &staffing.Department{
		ID:   id,
		Name: name,
	}

	return nil
}

func (r *departmentRepository) DeleteDepartment(ctx context.Context, departmentID staffing.DepartmentID) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.departments[departmentID] = nil
	return nil
}
