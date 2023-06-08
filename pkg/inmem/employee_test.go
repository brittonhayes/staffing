package inmem

import (
	"testing"

	"github.com/brittonhayes/staffing"
	"github.com/stretchr/testify/assert"
)

func Test_NewEmployeeRepository(t *testing.T) {
	got := NewEmployeeRepository()
	assert.Implements(t, (*staffing.EmployeeRepository)(nil), got)
}
