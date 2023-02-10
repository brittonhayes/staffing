package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/brittonhayes/staffing/inmem"
	"github.com/brittonhayes/staffing/project"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/brittonhayes/staffing/server"
	"google.golang.org/protobuf/proto"
)

func main() {

	projects := inmem.NewProjectRepository()
	logger := watermill.NewStdLogger(true, false)

	subscriber := gochannel.NewGoChannel(gochannel.Config{}, logger)
	publisher := gochannel.NewGoChannel(gochannel.Config{}, logger)

	ps := project.NewService(projects)
	srv := server.New(ps, publisher, subscriber, logger)

	_, err := subscriber.Subscribe(context.Background(), server.CreateProjectSubscribeTopic)
	if err != nil {
		panic(err)
	}

	go publishMessages(subscriber)

	err = srv.RunPubSub(context.Background())
	if err != nil {
		panic(err)
	}
}

// publish messages simulates a client publishing messages to the server
func publishMessages(publisher message.Publisher) {
	for {
		b, err := proto.Marshal(&pb.ProjectCreateCommand{
			Name: gofakeit.Name(),
		})
		if err != nil {
			panic(err)
		}

		msg := message.NewMessage(watermill.NewUUID(), b)
		if err := publisher.Publish(server.CreateProjectSubscribeTopic, msg); err != nil {
			panic(err)
		}

		time.Sleep(250 * time.Millisecond)
	}
}
