package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/pkg/project"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/go-chi/chi"
)

const (
	// CreateProjectHandlerName is the name of the handler for creating a project
	CreateProjectHandlerName = "create_project"
	// CreateProjectSubscribeTopic is the subscriber topic for creating a project
	CreateProjectSubscribeTopic = "topic.create_project"
	// CreateProjectPublishTopic is the topic for publishing a project created event
	CreateProjectPublishTopic = "topic.project_created"
)

type projectHttpHandler struct {
	service project.Service

	logger watermill.LoggerAdapter
}

func (h *projectHttpHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", h.createProjectHandler)
		r.Delete("/", h.cancelProjectCommand)
	})

	return r
}

func (h *projectHttpHandler) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.ProjectCreateCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	_, err = h.service.CreateProject(ctx, &request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *projectHttpHandler) cancelProjectCommand(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request pb.ProjectCancelCommand

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	_, err = h.service.CancelProject(ctx, &request)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type projectPubsubHandler struct {
	service project.Service

	logger watermill.LoggerAdapter
}

func (h *projectPubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddHandler(CreateProjectHandlerName, CreateProjectSubscribeTopic, subscriber, CreateProjectPublishTopic, publisher, h.createProject)
}

func (h *projectPubsubHandler) createProject(msg *message.Message) ([]*message.Message, error) {

	var command pb.ProjectCreateCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	_, err = h.service.CreateProject(context.Background(), &command)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
