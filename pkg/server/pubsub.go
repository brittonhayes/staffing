package server

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/brittonhayes/staffing/pkg/department"
	"github.com/brittonhayes/staffing/pkg/project"
)

type pubsubServer struct {
	Project    project.Service
	Department department.Service

	Logger watermill.LoggerAdapter

	router *message.Router
}

func NewPubSubServer(ps project.Service, ds department.Service, publisher message.Publisher, subscriber message.Subscriber, logger watermill.LoggerAdapter) Server {
	s := &pubsubServer{
		Project:    ps,
		Department: ds,
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
		Logger:          s.Logger,
	}.Middleware)

	ph := &projectPubsubHandler{
		service: ps,
		logger:  logger,
	}
	ph.addHandlers(router, publisher, subscriber)

	dh := &departmentPubsubHandler{
		service: ds,
		logger:  logger,
	}
	dh.addHandlers(router, publisher, subscriber)

	s.router = router

	return s
}

// RunPubSub runs the pubsub router
func (s *pubsubServer) Run(ctx context.Context) error {
	return s.router.Run(ctx)
}
