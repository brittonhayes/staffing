package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/department"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/proto/pb"
)

const (
	// CreateDepartmentHandlerName is the name of the handler for creating a department
	CreateDepartmentHandlerName = "create_department"
	// CreateDepartmentSubscribeTopic is the subscriber topic for creating a department
	CreateDepartmentSubscribeTopic = "topic.create_department"
	// CreateDepartmentPublishTopic is the topic for publishing a department created event
	CreateDepartmentPublishTopic = "topic.department_created"
)

type departmentPubSubHandler struct {
	service department.Service

	router *message.Router

	subscriber message.Subscriber
	publisher  message.Publisher

	logger watermill.LoggerAdapter
}

func (h *departmentPubSubHandler) addHandlers() {
	h.router.AddHandler(CreateDepartmentHandlerName, CreateDepartmentSubscribeTopic, h.subscriber, CreateDepartmentPublishTopic, h.publisher, h.createDepartment)
}

func (h *departmentPubSubHandler) createDepartment(msg *message.Message) ([]*message.Message, error) {

	var command pb.DepartmentCreateCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	h.logger.Debug("Received create department command", watermill.LogFields{"uuid": msg.UUID, "name": command.Name})

	err = h.service.CreateDepartment(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
