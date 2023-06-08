package inmem

import (
	"testing"

	"github.com/brittonhayes/staffing"
	"github.com/stretchr/testify/assert"
)

func Test_NewDepartmentRepository(t *testing.T) {
	got := NewDepartmentRepository()
	assert.Implements(t, (*staffing.DepartmentRepository)(nil), got)
}
