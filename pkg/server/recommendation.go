package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/pkg/recommend"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/pkg/errors"
)

const (
	CreateUserHandlerName    = "create_user"
	CreateUserSubscribeTopic = "topic.create_recommendation_user"
)

type recommendationPubsubHandler struct {
	service recommend.Service

	logger watermill.LoggerAdapter
}

func (h *recommendationPubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddNoPublisherHandler(CreateUserHandlerName, CreateEmployeePublishTopic, subscriber, h.createUser)
}

func (h *recommendationPubsubHandler) createUser(msg *message.Message) error {

	var command pb.EmployeeCreateCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal recommendation create user command")
	}

	createUser := &pb.RecommendationCreateUserCommand{
		UserId: command.GetName(),
		Labels: command.GetLabels(),
	}

	err = h.service.CreateUser(context.Background(), createUser)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}
