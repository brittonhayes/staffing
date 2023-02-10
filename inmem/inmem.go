package inmem

import (
	"context"
	"sync"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
)

type projectRepository struct {
	mu       sync.RWMutex
	projects map[staffing.ProjectID]*staffing.Project
}

func NewProjectRepository() staffing.ProjectRepository {
	return &projectRepository{
		projects: make(map[staffing.ProjectID]*staffing.Project),
	}
}

func (r *projectRepository) CreateProject(ctx context.Context, name string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.projects[staffing.ProjectID(name)] = &staffing.Project{
		ID:   staffing.ProjectID(uuid.NewString()),
		Name: name,
	}

	return nil
}

func (r *projectRepository) AssignEmployee(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *projectRepository) UnassignEmployee(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}
