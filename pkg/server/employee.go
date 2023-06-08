package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/pkg/employee"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/go-chi/chi"
)

const (
	// CreateEmployeeHandlerName is the name of the handler for creating a employee
	CreateEmployeeHandlerName = "create_employee"
	// CreateEmployeeSubscribeTopic is the subscriber topic for creating a employee
	CreateEmployeeSubscribeTopic = "topic.create_employee"
	// CreateEmployeePublishTopic is the topic for publishing a employee created event
	CreateEmployeePublishTopic = "topic.employee_created"
)

type employeeHttpHandler struct {
	service employee.Service

	logger watermill.LoggerAdapter
}

func (h *employeeHttpHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", h.createEmployeeHandler)
	})

	return r
}

func (h *employeeHttpHandler) createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.EmployeeCreateCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	err = h.service.CreateEmployee(ctx, &request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type employeePubsubHandler struct {
	service employee.Service

	logger watermill.LoggerAdapter
}

func (h *employeePubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddHandler(CreateEmployeeHandlerName, CreateEmployeeSubscribeTopic, subscriber, CreateProjectPublishTopic, publisher, h.createEmployee)
}

func (h *employeePubsubHandler) createEmployee(msg *message.Message) ([]*message.Message, error) {

	var command pb.EmployeeCreateCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	err = h.service.CreateEmployee(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
