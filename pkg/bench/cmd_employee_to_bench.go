package bench

import (
	"context"
	
	"github.com/ThreeDotsLabs/watermill"	
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
)

type MoveEmployeeToBenchHandler struct {
	eventBus *cqrs.EventBus
}

func NewMoveEmployeeToBenchHandler(eventBus *cqrs.EventBus) *MoveEmployeeToBenchHandler {
	return &MoveEmployeeToBenchHandler{
		eventBus: eventBus,
	}
}

func (m MoveEmployeeToBenchHandler) HandlerName() string {
	return "MoveEmployeeToBenchHandler"
}

func (m MoveEmployeeToBenchHandler) NewCommand() interface{} {
	return &pb.MoveEmployeeToBenchCommand{}
}

func (m MoveEmployeeToBenchHandler) Handle(ctx context.Context, c interface{}) error {
	cmd := c.(*pb.MoveEmployeeToBenchCommand)

	if err := m.eventBus.Publish(ctx, &pb.EmployeeMovedToBenchEvent{
		EmployeeId:   cmd.EmployeeId,
		EmployeeName: cmd.EmployeeName,
	}); err != nil {
		return err
	}

	
	logger.Info("Employee moved to bench", watermill.LogFields{
		"handler":     m.HandlerName(),
		"employee_name": cmd.EmployeeName,
	})

	return nil
}
