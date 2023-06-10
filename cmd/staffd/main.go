package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/pkg/department"
	"github.com/brittonhayes/staffing/pkg/employee"
	"github.com/brittonhayes/staffing/pkg/project"
	"github.com/brittonhayes/staffing/pkg/server"
	"github.com/brittonhayes/staffing/pkg/sqlite"
	"github.com/brittonhayes/staffing/proto/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	var (
		storage     = flag.String("storage", "inmemory", "select storage type from inmemory,sqlite (default inmemory)")
		debug       = flag.Bool("debug", false, "enable debug logging (default false)")
		trace       = flag.Bool("trace", false, "enable tracing (default false)")
		httpAddress = flag.String("address", ":8080", "HTTP address port (default :8080)")

		ctx = context.Background()
	)
	flag.Parse()
	logger := watermill.NewStdLogger(*debug, *trace)

	var (
		projects    staffing.ProjectRepository
		departments staffing.DepartmentRepository
		employees   staffing.EmployeeRepository
	)

	switch *storage {
	case "inmemory":
		logger.Debug("Using in-memory storage", nil)
		projects = sqlite.NewProjectRepository("file::memory:?cache=shared")
		defer projects.Close()

		departments = sqlite.NewDepartmentRepository("file::memory:?cache=shared")
		defer departments.Close()

		employees = sqlite.NewEmployeeRepository("file::memory:?cache=shared")
		defer employees.Close()
	case "sqlite":
		logger.Debug("Using sqlite storage", nil)
		projects = sqlite.NewProjectRepository("file:projects.db")
		defer projects.Close()

		departments = sqlite.NewDepartmentRepository("file:departments.db")
		defer departments.Close()

		employees = sqlite.NewEmployeeRepository("file:employees.db")
		defer employees.Close()
	}

	fieldKeys := []string{"method"}

	projectSvc := project.NewService(projects)
	projectSvc = project.NewLoggingService(logger.With(watermill.LogFields{"service": "project"}), projectSvc)
	projectSvc = project.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "project_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "project_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		projectSvc,
	)

	departmentSvc := department.NewService(departments)
	departmentSvc = department.NewLoggingService(logger.With(watermill.LogFields{"service": "department"}), departmentSvc)
	departmentSvc = department.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "department_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "department_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		departmentSvc,
	)

	employeeSvc := employee.NewService(employees)
	employeeSvc = employee.NewLoggingService(logger.With(watermill.LogFields{"service": "employee"}), employeeSvc)
	employeeSvc = employee.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "employee_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "employee_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		employeeSvc,
	)

	subscriber := gochannel.NewGoChannel(gochannel.Config{}, logger)
	publisher := gochannel.NewGoChannel(gochannel.Config{}, logger)

	httpServer := server.NewHTTPServer(projectSvc, departmentSvc, employeeSvc, *httpAddress, logger)
	pubsubServer := server.NewPubSubServer(projectSvc, departmentSvc, employeeSvc, subscriber, publisher, logger)

	// Run servers in multiple goroutines
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		log.Fatal(pubsubServer.Run(ctx))
		wg.Done()
	}()
	go func() {
		log.Fatal(httpServer.Run(ctx))
		wg.Done()
	}()

	// go publishMessages(ctx, subscriber)
	wg.Wait()
}

// publish messages simulates a client publishing messages to the server
func publishMessages(ctx context.Context, publisher message.Publisher) {

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

		time.Sleep(5 * time.Second)
	}
}
