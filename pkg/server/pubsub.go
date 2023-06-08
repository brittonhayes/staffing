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
)

type pubsubServer struct {
	Project    project.Service
	Department department.Service
	Employee   employee.Service

	Logger watermill.LoggerAdapter

	router *message.Router
}

func NewPubSubServer(projectService project.Service, departmentService department.Service, employeeService employee.Service, publisher message.Publisher, subscriber message.Subscriber, logger watermill.LoggerAdapter) Server {
	server := &pubsubServer{
		Project:    projectService,
		Department: departmentService,
		Employee:   employeeService,
		Logger:     logger,
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer, middleware.Retry{
		MaxRetries:      3,
		InitialInterval: time.Millisecond * 100,
		Logger:          server.Logger,
	}.Middleware)

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

	server.router = router

	return server
}

// RunPubSub runs the pubsub router
func (s *pubsubServer) Run(ctx context.Context) error {
	return s.router.Run(ctx)
}
