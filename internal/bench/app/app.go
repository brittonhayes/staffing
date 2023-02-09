package app

import (
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app/command"
)

type Application struct {
	Commands Commands
}

type Commands struct {
	MoveEmployeeToBench    command.MoveEmployeeToBenchHandler
	AlignEmployeeToProject command.AlignEmployeeToProjectHandler
}
