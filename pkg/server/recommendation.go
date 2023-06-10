package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing/internal/protobuf"
	"github.com/brittonhayes/staffing/pkg/recommend"
	"github.com/brittonhayes/staffing/proto/pb"
)

const (
	CreateUserHandlerName    = "create_user"
	CreateUserSubscribeTopic = "topic.create_recommendation_user"
	CreateUserPublishTopic   = "topic.recommendation_user_created"
)

type recommendationPubsubHandler struct {
	service recommend.Service

	logger watermill.LoggerAdapter
}

func (h *recommendationPubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddHandler(CreateUserHandlerName, CreateEmployeePublishTopic, subscriber, CreateUserPublishTopic, publisher, h.createRecommendation)
}

func (h *recommendationPubsubHandler) createRecommendation(msg *message.Message) ([]*message.Message, error) {

	var command pb.RecommendationCreateUserCommand
	p := protobuf.ProtobufMarshaler{}

	err := p.Unmarshal(msg, &command)
	if err != nil {
		return nil, err
	}

	err = h.service.CreateUser(context.Background(), &command)
	if err != nil {
		h.logger.Error("error", err, nil)
		return nil, err
	}

	h.logger.Debug("CreateUser command received", nil)

	return nil, nil
}
