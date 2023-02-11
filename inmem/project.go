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
	repository := &projectRepository{
		projects: make(map[staffing.ProjectID]*staffing.Project),
	}

	repository.projects[cloudMigration.ID] = cloudMigration
	repository.projects[marketing.ID] = marketing

	return repository
}

func (r *projectRepository) Close() error {
	defer r.mu.Unlock()
	return nil
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

	if r.projects[projectID] == nil {
		return staffing.ErrProjectNotFound
	}

	r.projects[projectID].AssignedEmployees = append(r.projects[projectID].AssignedEmployees, employeeID)

	return nil
}

func (r *projectRepository) UnassignEmployee(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.projects[projectID] == nil {
		return staffing.ErrProjectNotFound
	}

	for i, id := range r.projects[projectID].AssignedEmployees {
		if id == employeeID {
			r.projects[projectID].AssignedEmployees = append(r.projects[projectID].AssignedEmployees[:i], r.projects[projectID].AssignedEmployees[i+1:]...)
			break
		}
	}

	return nil
}
