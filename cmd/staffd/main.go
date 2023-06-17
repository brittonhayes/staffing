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
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/pkg/department"
	"github.com/brittonhayes/staffing/pkg/employee"
	"github.com/brittonhayes/staffing/pkg/project"
	"github.com/brittonhayes/staffing/pkg/recommend"
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
		amqpURI     = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "RabbitMQ PubSub URI (default amqp://guest:guest@rabbitmq:5672/)")

		ctx = context.Background()
	)

	flag.Parse()
	logger := watermill.NewStdLogger(*debug, *trace)

	var (
		projects        staffing.ProjectRepository
		departments     staffing.DepartmentRepository
		employees       staffing.EmployeeRepository
		recommendations staffing.RecommendationRepository
	)

	switch *storage {
	case "inmemory":
		logger.Debug("Using in-memory storage", nil)
		projects = sqlite.NewProjectRepository("file::memory:?cache=shared", true)
		defer projects.Close()

		departments = sqlite.NewDepartmentRepository("file::memory:?cache=shared", true)
		defer departments.Close()

		employees = sqlite.NewEmployeeRepository("file::memory:?cache=shared", true)
		defer employees.Close()

		recommendations = sqlite.NewRecommendationRepository("file::memory:?cache=shared", true)
		defer recommendations.Close()

	case "sqlite":
		logger.Debug("Using sqlite storage", nil)
		projects = sqlite.NewProjectRepository("file:projects.db", false)
		defer projects.Close()

		departments = sqlite.NewDepartmentRepository("file:departments.db", false)
		defer departments.Close()

		employees = sqlite.NewEmployeeRepository("file:employees.db", false)
		defer employees.Close()

		recommendations = sqlite.NewRecommendationRepository("file:recommendations.db", false)
		defer recommendations.Close()
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

	recommendationSvc := recommend.NewService(recommendations)
	recommendationSvc = recommend.NewLoggingService(logger.With(watermill.LogFields{"service": "recommendation"}), recommendationSvc)
	recommendationSvc = recommend.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "recommendation_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "recommendation_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		recommendationSvc,
	)

	// pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	amqpPubConfig := amqp.NewDurablePubSubConfig(*amqpURI, nil)
	amqpSubConfig := amqp.NewDurableQueueConfig(*amqpURI)

	publisher, err := amqp.NewPublisher(amqpPubConfig, logger)
	if err != nil {
		log.Fatal(err)
	}

	subscriber, err := amqp.NewSubscriber(amqpSubConfig, logger)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := server.NewHTTPServer(projectSvc, departmentSvc, employeeSvc, *httpAddress, logger)
	pubsubServer := server.NewPubSubServer(projectSvc, departmentSvc, employeeSvc, recommendationSvc, publisher, subscriber, logger)

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

	go publishMessages(ctx, publisher, 3)
	wg.Wait()
}

// publish messages simulates a client publishing messages to the server
func publishMessages(ctx context.Context, publisher message.Publisher, count int) {

	for i := 0; i <= count; i++ {
		employeeCmd, err := proto.Marshal(&pb.EmployeeCreateCommand{
			Name:   gofakeit.Name(),
			Labels: []string{"language:go"},
		})
		if err != nil {
			panic(err)
		}

		err = publisher.Publish(server.CreateEmployeeTopic, message.NewMessage(watermill.NewUUID(), employeeCmd))
		if err != nil {
			panic(err)
		}

		time.Sleep(3 * time.Second)
	}

}
