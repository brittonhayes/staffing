package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/department"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/go-chi/chi"
)

const (
	// CreateDepartmentHandlerName is the name of the handler for creating a department
	CreateDepartmentHandlerName = "create_department"
	// CreateDepartmentSubscribeTopic is the subscriber topic for creating a department
	CreateDepartmentSubscribeTopic = "topic.create_department"
	// CreateDepartmentPublishTopic is the topic for publishing a department created event
	CreateDepartmentPublishTopic = "topic.department_created"
)

type departmentHandler struct {
	service department.Service

	logger watermill.LoggerAdapter
}

func (h *departmentHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/department", func(r chi.Router) {
		r.Post("/", h.createDepartmentHandler)
	})

	return r
}

func (h *departmentHandler) addPubsubHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddHandler(CreateDepartmentHandlerName, CreateDepartmentSubscribeTopic, subscriber, CreateDepartmentPublishTopic, publisher, h.createDepartment)
}

func (h *departmentHandler) createDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.DepartmentCreateCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	err = h.service.CreateDepartment(ctx, &request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *departmentHandler) createDepartment(msg *message.Message) ([]*message.Message, error) {

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
