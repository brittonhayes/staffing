package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/ports"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
	"github.com/brittonhayes/go-cqrs-example/pkg/logging"
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logging.New(false, false, "router"))
	if err != nil {
		panic(err)
	}

	service := ports.NewService(app.Application{}, router)

	go publishCommands(service.CommandBus())

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func publishCommands(commandBus *cqrs.CommandBus) func() {
	i := 0
	for {
		i++

		moveEmployeeToBenchCmd := &pb.MoveEmployeeToBenchCommand{
			EmployeeId:   gofakeit.UUID(),
			EmployeeName: gofakeit.Name(),
		}
		if err := commandBus.Send(context.Background(), moveEmployeeToBenchCmd); err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)
	}
}
