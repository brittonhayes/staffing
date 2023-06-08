package staffing

import (
	"context"
	"errors"
)

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	CreateProject(ctx context.Context, name string) error
	DeleteProject(ctx context.Context, projectID ProjectID) error
	Close() error
}

type ProjectID string

type Project struct {
	ID   ProjectID
	Name string
}
