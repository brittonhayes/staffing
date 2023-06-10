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

func (r *projectRepository) CreateProject(ctx context.Context, name string) (*staffing.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id := staffing.ProjectID(uuid.NewString())
	r.projects[id] = &staffing.Project{
		ID:   id,
		Name: name,
	}

	return r.projects[id], nil
}

func (r *projectRepository) CancelProject(ctx context.Context, projectID staffing.ProjectID) (*staffing.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.projects[projectID] == nil {
		return nil, staffing.ErrProjectNotFound
	}

	r.projects[projectID] = nil

	return &staffing.Project{ID: projectID}, nil
}
