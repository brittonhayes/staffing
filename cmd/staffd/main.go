package main

import (
	"context"
	"flag"
	"time"

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

	subscriber := gochannel.NewGoChannel(gochannel.Config{}, logger)
	publisher := gochannel.NewGoChannel(gochannel.Config{}, logger)

	ps := project.NewService(projects)
	ds := department.NewService(departments)
	srv := server.New(ps, ds, publisher, subscriber, logger)

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
