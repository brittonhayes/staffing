package bench

import (
	"context"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
)

type EmployeesBenchReport struct {
	employees map[string]struct{}
	mu        sync.Mutex
}

func NewEmployeesBenchReport() *EmployeesBenchReport {
	return &EmployeesBenchReport{
		employees: make(map[string]struct{}),
	}
}

func (b EmployeesBenchReport) HandlerName() string {
	return "EmployeesBenchReport"
}

func (b EmployeesBenchReport) NewEvent() interface{} {
	return &pb.EmployeeMovedToBenchEvent{}
}

func (b *EmployeesBenchReport) Handle(ctx context.Context, e interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	event := e.(*pb.EmployeeMovedToBenchEvent)

	if _, ok := b.employees[event.EmployeeId]; ok {
		return nil
	}

	b.employees[event.EmployeeId] = struct{}{}
	logger.Info("Employee Bench Report", watermill.LogFields{
		"handler":                 b.HandlerName(),
		"benched_employees_count": len(b.employees),
	})

	return nil
}
