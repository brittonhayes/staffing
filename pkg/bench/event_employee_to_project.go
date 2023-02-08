package bench

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
)

type EmployeeAlignedToProject struct {
	commandBus *cqrs.CommandBus
}

func NewEmployeeAlignedToProjectHandler(commandBus *cqrs.CommandBus) *EmployeeAlignedToProject {
	return &EmployeeAlignedToProject{
		commandBus: commandBus,
	}
}

func (a EmployeeAlignedToProject) HandlerName() string {
	return "EmployeeAlignedToProject"
}

func (a EmployeeAlignedToProject) NewEvent() interface{} {
	return &pb.EmployeeMovedToBenchEvent{}
}

func (a EmployeeAlignedToProject) Handle(ctx context.Context, e interface{}) error {
	event := e.(*pb.EmployeeMovedToBenchEvent)

	if err := a.commandBus.Send(ctx, &pb.AlignEmployeeToProjectCommand{
		EmployeeId:   event.EmployeeId,
		EmployeeName: event.EmployeeName,
		ProjectId:    gofakeit.UUID(),
		ProjectName:  gofakeit.Company(),
	}); err != nil {
		return err
	}

	logger.Info("Employee aligned to project from bench", watermill.LogFields{
		"handler":     a.HandlerName(),
		"employee_id": event.EmployeeId,
	})

	return nil
}
