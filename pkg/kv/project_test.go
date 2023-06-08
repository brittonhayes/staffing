package kv

import (
	"os"
	"testing"

	"github.com/brittonhayes/staffing"
	"github.com/stretchr/testify/assert"
)

func Test_NewProjectRepository(t *testing.T) {

	f, err := os.CreateTemp("", "*.db")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	got := NewProjectRepository(f.Name())
	assert.Implements(t, (*staffing.ProjectRepository)(nil), got)
}
