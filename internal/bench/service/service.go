package service

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app/command"
)

func NewApplication(ctx context.Context, eb *cqrs.EventBus, cb *cqrs.CommandBus) app.Application {
	return newApplication(ctx, eb, cb)
}

func newApplication(ctx context.Context, eb *cqrs.EventBus, cb *cqrs.CommandBus) app.Application {
	return app.Application{
		Commands: app.Commands{
			MoveEmployeeToBench:    command.NewMoveEmployeeToBenchHandler(eb),
			AlignEmployeeToProject: command.NewAlignEmployeeToProjectHandler(eb),
		},
	}
}
