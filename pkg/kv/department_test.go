package kv

import (
	"os"
	"testing"

	"github.com/brittonhayes/staffing"
	"github.com/stretchr/testify/assert"
)

func Test_NewDepartmentRepository(t *testing.T) {

	f, err := os.CreateTemp("", "*.db")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	got := NewDepartmentRepository(f.Name())
	assert.Implements(t, (*staffing.DepartmentRepository)(nil), got)
}
