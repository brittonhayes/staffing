package ports

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app"
	"github.com/brittonhayes/go-cqrs-example/internal/bench/app/command"
	"github.com/brittonhayes/go-cqrs-example/pkg/logging"
	"github.com/brittonhayes/go-cqrs-example/pkg/protobuf"
	"github.com/brittonhayes/go-cqrs-example/pkg/pubsub"
)

var (
	logger     = logging.New(false, false, "cqrs")
	brokers    = []string{"kafka:9092"}
	marshaller = kafka.DefaultMarshaler{}
)

type CqrsService struct {
	cqrs *cqrs.Facade
	app  app.Application
}

func NewService(app app.Application, router *message.Router) CqrsService {
	return CqrsService{
		cqrs: createCQRS(router),
		app:  app,
	}
}

func (s *CqrsService) CommandBus() *cqrs.CommandBus {
	return s.cqrs.CommandBus()
}

func createCQRS(router *message.Router) *cqrs.Facade {
	subscriberConfig := kafka.SubscriberConfig{
		Brokers:     []string{"kafka:9092"},
		Unmarshaler: kafka.DefaultMarshaler{},
	}

	fascade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		CommandEventMarshaler: protobuf.ProtobufMarshaler{},
		Router:                router,
		Logger:                logger,
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
		CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
			return []cqrs.CommandHandler{
				command.NewMoveEmployeeToBenchHandler(eb),
				command.NewAlignEmployeeToProjectHandler(eb),
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
				command.NewEmployeesBenchReport(),
				command.NewEmployeeAlignedToProjectHandler(cb),
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

	return fascade
}
