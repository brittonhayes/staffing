package server

import (
	"context"
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

	var employee *staffing.Employee

	// WARN: This is likely where the panic
	// "invalid memory address or nil pointer dereference" is occurring
	if err := json.Unmarshal(msg.Payload, &employee); err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}

		// FIX: Recommendation from Milosz to ack bad message until good comes through
		// for debug
		// defer msg.Ack()

		return errors.Wrap(err, "failed to unmarshal json from staffing.Employee")
	}

	log.Printf("successfully unmarshalled msg: %v", employee)
	user := &pb.RecommendationCreateUserCommand{
		UserId: string(employee.ID),
	}

	err := h.service.CreateUser(context.Background(), user)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}
