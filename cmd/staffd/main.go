package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/go-cqrs-example/pkg/bench"
	"github.com/brittonhayes/go-cqrs-example/pkg/grpc/pb"
	"github.com/brittonhayes/go-cqrs-example/pkg/protobuf"
	"github.com/brittonhayes/go-cqrs-example/pkg/pubsub"
)

var (
	brokers    = []string{"kafka:9092"}
	marshaller = kafka.DefaultMarshaler{}
	timeout    = time.Second * 5
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	subscriberConfig := kafka.SubscriberConfig{
		Brokers:     []string{"kafka:9092"},
		Unmarshaler: kafka.DefaultMarshaler{},
	}

	// CQRS is built on messages router. Detailed documentation: https://watermill.io/docs/messages-router/
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer, middleware.Timeout(timeout))

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		CommandEventMarshaler: protobuf.ProtobufMarshaler{},
		Router:                router,
		Logger:                logger,
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
		CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
			return []cqrs.CommandHandler{
				bench.NewMoveEmployeeToBenchHandler(eb),
				bench.NewAlignEmployeeToProjectHandler(eb),
			}
		},
		CommandsPublisher: pubsub.NewPublisher(kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: marshaller,
		}, logger),
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return kafka.NewSubscriber(subscriberConfig, logger)
		},
		GenerateEventsTopic: func(eventName string) string {
			return eventName
		},
		EventHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
			return []cqrs.EventHandler{
				bench.NewEmployeesBenchReport(),
				bench.NewEmployeeAlignedToProjectHandler(cb),
			}
		},
		EventsPublisher: pubsub.NewPublisher(kafka.PublisherConfig{Brokers: brokers, Marshaler: marshaller}, logger),
		EventsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return kafka.NewSubscriber(subscriberConfig, logger)
		},
	})
	if err != nil {
		panic(err)
	}

	go publishCommands(cqrsFacade.CommandBus())

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

		time.Sleep(1 * time.Second)
	}
}
