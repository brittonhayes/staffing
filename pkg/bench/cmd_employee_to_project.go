package bench

import (
	"context"
	"math/rand"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
	"github.com/pkg/errors"
)

type AlignEmployeeToProjectHandler struct {
	projects map[string]struct{}
	eventBus *cqrs.EventBus

	mu sync.Mutex
}

func NewAlignEmployeeToProjectHandler(eventBus *cqrs.EventBus) *AlignEmployeeToProjectHandler {
	return &AlignEmployeeToProjectHandler{
		projects: make(map[string]struct{}),
		eventBus: eventBus,
	}
}

func (a AlignEmployeeToProjectHandler) HandlerName() string {
	return "AlignEmployeeToProjectHandler"
}

func (a AlignEmployeeToProjectHandler) NewCommand() interface{} {
	return &pb.AlignEmployeeToProjectCommand{}
}

func (a AlignEmployeeToProjectHandler) Handle(ctx context.Context, c interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	cmd := c.(*pb.AlignEmployeeToProjectCommand)

	if rand.Int63n(10) == 0 {
		return errors.Errorf("no positions open on project %s", cmd.ProjectId)
	}

	if _, ok := a.projects[cmd.ProjectId]; ok {
		return nil
	}

	a.projects[cmd.ProjectId] = struct{}{}

	if err := a.eventBus.Publish(ctx, &pb.EmployeeAlignedToProjectEvent{
		EmployeeId: cmd.EmployeeId,
	}); err != nil {
		return err
	}

	logger.Info("Employee aligned to project", watermill.LogFields{
		"handler":        a.HandlerName(),
		"employee_id":    cmd.EmployeeId,
		"project_id":     cmd.ProjectId,
		"projects_count": len(a.projects),
	})

	return nil
}
