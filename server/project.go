package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/project"
	"github.com/brittonhayes/staffing/proto/pb"
)

const (
	// CreateProjectHandlerName is the name of the handler for creating a project
	CreateProjectHandlerName = "create_project"
	// CreateProjectSubscribeTopic is the subscriber topic for creating a project
	CreateProjectSubscribeTopic = "topic.create_project"
	// CreateProjectPublishTopic is the topic for publishing a create project command
	CreateProjectPublishTopic = "topic.project_created"
)

type projectHandler struct {
	service project.Service

	router *message.Router

	subscriber message.Subscriber
	publisher  message.Publisher

	logger watermill.LoggerAdapter
}

func (h *projectHandler) addHandlers() {
	h.router.AddHandler(CreateProjectHandlerName, CreateProjectSubscribeTopic, h.subscriber, CreateProjectPublishTopic, h.publisher, h.createProject)
}

func (h *projectHandler) createProject(msg *message.Message) ([]*message.Message, error) {

	var command pb.ProjectCreateCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	h.logger.Info("Received create project command", watermill.LogFields{"uuid": msg.UUID, "name": command.Name})

	err = h.service.CreateProject(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
