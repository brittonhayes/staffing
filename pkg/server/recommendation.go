package server

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/pkg/recommend"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/pkg/errors"
)

const (
	CreateUserHandlerName = "create_user"
)

type recommendationPubsubHandler struct {
	service recommend.Service

	logger watermill.LoggerAdapter
}

func (h *recommendationPubsubHandler) addHandlers(router *message.Router, publisher message.Publisher, subscriber message.Subscriber) {
	router.AddNoPublisherHandler(CreateUserHandlerName, EmployeeCreatedTopic, subscriber, h.createUser)
}

func (h *recommendationPubsubHandler) createUser(msg *message.Message) error {

	defer msg.Ack()

	h.logger.Debug("received message", watermill.LogFields{
		"msg": msg.UUID,
	})

	var employee staffing.Employee
	if err := json.Unmarshal(msg.Payload, &employee); err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}

		return errors.Wrap(err, "failed to unmarshal json from staffing.Employee")
	}
	user := pb.RecommendationCreateUserCommand{
		UserId: string(employee.ID),
	}

	err := h.service.CreateUser(msg.Context(), &user)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}
