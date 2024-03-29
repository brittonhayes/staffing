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
	// CreateEmployeeTopic is the subscriber topic for creating a employee
	CreateEmployeeTopic = "topic.create_employee"
	// EmployeeCreatedTopic is the topic for publishing a employee created event
	EmployeeCreatedTopic = "topic.employee_created"

	// DeleteEmployeeHandlerName is the name of the handler for deleting a employee
	DeleteEmployeeHandlerName = "delete_employee"
	// DeleteEmployeeTopic is the subscriber topic for deleting a employee
	DeleteEmployeeTopic = "topic.delete_employee"
	// EmployeeDeletedTopic is the topic for publishing a employee deleted event
	EmployeeDeletedTopic = "topic.employee_deleted"
)

type employeeHttpHandler struct {
	service employee.Service

	logger watermill.LoggerAdapter
}

func (h *employeeHttpHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", h.createEmployeeHandler)
		r.Delete("/", h.deleteEmployeeHandler)
	})

	return r
}

func (h *employeeHttpHandler) deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.EmployeeDeleteCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	err = h.service.DeleteEmployee(ctx, &request)
	if err != nil {
		h.logger.Error("error", err, nil)
		encodeError(ctx, err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&request)
}

func (h *employeeHttpHandler) createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.EmployeeCreateCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	resp, err := h.service.CreateEmployee(ctx, &request)
	if err != nil {
		h.logger.Error("error", err, nil)
		encodeError(ctx, err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

type employeePubsubHandler struct {
	service employee.Service

	logger watermill.LoggerAdapter
}

func (h *employeePubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddHandler(CreateEmployeeHandlerName, CreateEmployeeTopic, subscriber, EmployeeCreatedTopic, publisher, h.createEmployee)
	router.AddHandler(DeleteEmployeeHandlerName, DeleteEmployeeTopic, subscriber, EmployeeDeletedTopic, publisher, h.deleteEmployee)
}

func (h *employeePubsubHandler) createEmployee(msg *message.Message) ([]*message.Message, error) {

	var command pb.EmployeeCreateCommand
	p := protobuf.ProtobufMarshaler{}

	if err := p.Unmarshal(msg, &command); err != nil {
		return nil, err
	}

	resp, err := h.service.CreateEmployee(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return []*message.Message{
		message.NewMessage(watermill.NewUUID(), payload),
	}, nil

}

func (h *employeePubsubHandler) deleteEmployee(msg *message.Message) ([]*message.Message, error) {

	var command pb.EmployeeDeleteCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	err = h.service.DeleteEmployee(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
