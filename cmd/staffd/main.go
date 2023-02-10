package main

import (
	"context"
	"flag"
	"time"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/department"
	"github.com/brittonhayes/staffing/inmem"
	"github.com/brittonhayes/staffing/kv"
	"github.com/brittonhayes/staffing/project"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/brittonhayes/staffing/server"
	"google.golang.org/protobuf/proto"
)

func main() {
	var (
		inmemory = flag.Bool("inmem", false, "use in-memory repositories")
		debug    = flag.Bool("debug", false, "enable debug logging")

		ctx = context.Background()
	)
	flag.Parse()
	logger := watermill.NewStdLogger(*debug, false)

	var (
		projects    staffing.ProjectRepository
		departments staffing.DepartmentRepository
	)

	if *inmemory {
		logger.Debug("Using in-memory repositories", nil)
		projects = inmem.NewProjectRepository()
		departments = inmem.NewDepartmentRepository()
	} else {
		projects = kv.NewProjectRepository("projects.db")
		departments = kv.NewDepartmentRepository("departments.db")
		defer projects.Close()
		defer departments.Close()
	}

	fieldKeys := []string{"method"}

	subscriber := gochannel.NewGoChannel(gochannel.Config{}, logger)
	publisher := gochannel.NewGoChannel(gochannel.Config{}, logger)

	projectSvc := project.NewService(projects)
	projectSvc = project.NewLoggingService(logger.With(watermill.LogFields{"service": "project"}), projectSvc)
	projectSvc = project.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "pubsub",
		Subsystem: "project_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "pubsub",
			Subsystem: "project_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		projectSvc,
	)

	departmentSvc := department.NewService(departments)
	departmentSvc = department.NewLoggingService(logger.With(watermill.LogFields{"service": "department"}), departmentSvc)
	departmentSvc = department.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "pubsub",
		Subsystem: "department_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "pubsub",
			Subsystem: "department_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		departmentSvc,
	)

	srv := server.New(projectSvc, departmentSvc, publisher, subscriber, logger)
	_, err := subscriber.Subscribe(ctx, server.CreateProjectSubscribeTopic)
	if err != nil {
		panic(err)
	}

	go publishMessages(subscriber)

	err = srv.RunPubSub(ctx)
	if err != nil {
		panic(err)
	}
}

// publish messages simulates a client publishing messages to the server
func publishMessages(publisher message.Publisher) {
	for {
		projectCmd, err := proto.Marshal(&pb.ProjectCreateCommand{
			Name: gofakeit.Name(),
		})
		if err != nil {
			panic(err)
		}

		departmentCmd, err := proto.Marshal(&pb.DepartmentCreateCommand{
			Name: gofakeit.Verb(),
		})
		if err != nil {
			panic(err)
		}

		if err := publisher.Publish(server.CreateProjectSubscribeTopic, message.NewMessage(watermill.NewUUID(), projectCmd)); err != nil {
			panic(err)
		}

		if err := publisher.Publish(server.CreateDepartmentSubscribeTopic, message.NewMessage(watermill.NewUUID(), departmentCmd)); err != nil {
			panic(err)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
