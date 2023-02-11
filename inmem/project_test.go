package inmem

import (
	"context"
	"testing"

	"github.com/brittonhayes/staffing"
	"github.com/stretchr/testify/assert"
)

func Test_NewProjectRepository(t *testing.T) {
	got := NewProjectRepository()
	assert.Implements(t, (*staffing.ProjectRepository)(nil), got)
}

func Test_ProjectRepository_CreateProject(t *testing.T) {
	repo := NewProjectRepository()
	err := repo.CreateProject(context.Background(), marketing.Name)

	assert.NoError(t, err)
}

func Test_ProjectRepository_AssignEmployee(t *testing.T) {
	repo := NewProjectRepository()
	err := repo.AssignEmployee(context.Background(), marketing.ID, john.ID)

	assert.NoError(t, err)
}
