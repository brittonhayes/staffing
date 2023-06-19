package server

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/brittonhayes/staffing/pkg/department"
	"github.com/brittonhayes/staffing/pkg/employee"
	"github.com/brittonhayes/staffing/pkg/project"
	"github.com/brittonhayes/staffing/pkg/recommend"
)

const (
	DeadletterQueueTopic = "topic.deadletter_queue"
)

type pubsubServer struct {
	Project        project.Service
	Department     department.Service
	Employee       employee.Service
	Recommendation recommend.Service

	Logger watermill.LoggerAdapter

	router *message.Router
}

func NewPubSubServer(projectService project.Service, departmentService department.Service, employeeService employee.Service, recommendationService recommend.Service, publisher message.Publisher, subscriber message.Subscriber, logger watermill.LoggerAdapter) Server {
	server := &pubsubServer{
		Project:        projectService,
		Department:     departmentService,
		Employee:       employeeService,
		Recommendation: recommendationService,
		Logger:         logger,
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	poisonQueue, err := middleware.PoisonQueue(publisher, DeadletterQueueTopic)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(
		poisonQueue,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          server.Logger,
		}.Middleware,
		middleware.Recoverer,
	)

	router.AddPlugin(plugin.SignalsHandler)

	projectHandler := &projectPubsubHandler{
		service: projectService,
		logger:  logger,
	}
	projectHandler.addHandlers(router, publisher, subscriber)

	departmentHandler := &departmentPubsubHandler{
		service: departmentService,
		logger:  logger,
	}
	departmentHandler.addHandlers(router, publisher, subscriber)

	employeeHandler := &employeePubsubHandler{
		service: employeeService,
		logger:  logger,
	}
	employeeHandler.addHandlers(router, publisher, subscriber)

	recommendationHandler := &recommendationPubsubHandler{
		service: recommendationService,
		logger:  logger,
	}
	recommendationHandler.addHandlers(router, publisher, subscriber)

	server.router = router

	return server
}

// RunPubSub runs the pubsub router
func (s *pubsubServer) Run(ctx context.Context) error {
	return s.router.Run(ctx)
}
