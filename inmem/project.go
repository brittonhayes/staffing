package inmem

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/brittonhayes/staffing"
	"github.com/google/uuid"
)

type projectRepository struct {
	mu     sync.RWMutex
	events map[string]*staffing.Event
}

func NewProjectRepository() staffing.ProjectRepository {
	r := &projectRepository{
		events: make(map[string]*staffing.Event),
	}

	marketingEvent, _ := json.Marshal(marketing)
	r.events[uuid.NewString()] = &staffing.Event{
		ID:        uuid.NewString(),
		Timestamp: time.Now(),
		Type:      "project_created",
		Body:      marketingEvent,
	}

	return r
}

func (r *projectRepository) Close() error {
	defer r.mu.Unlock()
	return nil
}

func (r *projectRepository) CreateProject(ctx context.Context, name string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	payload := &staffing.Project{
		ID:   staffing.ProjectID(uuid.NewString()),
		Name: name,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	eventID := uuid.NewString()
	r.events[eventID] = &staffing.Event{
		ID:        eventID,
		Timestamp: time.Now(),
		Type:      "project_created",
		Body:      body,
	}

	for _, e := range r.events {
		if e.ID == eventID {
			var project staffing.Project
			json.Unmarshal(e.Body, &project)
			log.Println(project)
		}
	}

	return nil
}

func (r *projectRepository) AssignEmployee(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.NewString()
	r.events[id] = &staffing.Event{
		ID:        id,
		Timestamp: time.Now(),
		Type:      "employee_assigned",
		Body:      []byte(`{"project_id": "` + string(projectID) + `", "employee_id": "` + string(employeeID) + `"}`),
	}

	return nil
}

func (r *projectRepository) UnassignEmployee(ctx context.Context, projectID staffing.ProjectID, employeeID staffing.EmployeeID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.NewString()
	r.events[id] = &staffing.Event{
		ID:        id,
		Timestamp: time.Now(),
		Type:      "employee_unassigned",
		Body:      []byte(`{"project_id": "` + string(projectID) + `", "employee_id": "` + string(employeeID) + `"}`),
	}

	return nil
}
