package server

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/department"
	"github.com/brittonhayes/staffing/project"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// httpServer holds the dependencies for an HTTP server with pubsub handlers
type httpServer struct {
	Project    project.CommandService
	Department department.Service

	Logger watermill.LoggerAdapter

	router  chi.Router
	address string
}

// NewHTTPServer returns a new Server
func NewHTTPServer(ps project.CommandService, ds department.Service, address string, logger watermill.LoggerAdapter) Server {
	s := &httpServer{
		Project:    ps,
		Department: ds,
		Logger:     logger,
		address:    address,
	}

	ph := &projectHttpHandler{
		service: ps,
		logger:  logger,
	}

	dh := &departmentHttpHandler{
		service: ds,
		logger:  logger,
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/projects", ph.router())
		r.Mount("/departments", dh.router())
	})

	router.Method(http.MethodGet, "/metrics", promhttp.Handler())

	s.router = router

	return s
}

func (s *httpServer) Run(ctx context.Context) error {
	// The HTTP Server
	server := &http.Server{Addr: "0.0.0.0" + s.address, Handler: s.router}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(ctx)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				s.Logger.Error("graceful shutdown timed out.. forcing exit.", nil, nil)
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			s.Logger.Error("error shutting down server", err, nil)
			panic(err)
		}
		serverStopCtx()
	}()

	return http.ListenAndServe(s.address, s.router)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case staffing.ErrProjectNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
