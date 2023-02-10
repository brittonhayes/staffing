package server

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/brittonhayes/staffing/department"
	"github.com/brittonhayes/staffing/project"
)

type Server struct {
	Project project.Service

	Logger watermill.LoggerAdapter

	router *message.Router
}

func New(ps project.Service, ds department.Service, publisher message.Publisher, subscriber message.Subscriber, logger watermill.LoggerAdapter) *Server {
	s := &Server{
		Project: ps,
		Logger:  logger,
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

	ph := &projectPubSubHandler{
		service:    ps,
		router:     router,
		publisher:  publisher,
		subscriber: subscriber,
		logger:     logger,
	}
	ph.addHandlers()

	dh := &departmentPubSubHandler{
		service:    ds,
		router:     router,
		publisher:  publisher,
		subscriber: subscriber,
		logger:     logger,
	}
	dh.addHandlers()

	s.router = router

	return s
}

func (s *Server) RunPubSub(ctx context.Context) error {
	return s.router.Run(ctx)
}
