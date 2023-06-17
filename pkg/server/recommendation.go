package server

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/pkg/recommend"
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
	router.AddNoPublisherHandler(CreateUserHandlerName, CreateEmployeePublishTopic, subscriber, h.createUser)
}

func (h *recommendationPubsubHandler) createUser(msg *message.Message) error {

	var employee staffing.Employee

	// WARN: This is likely where the panic
	// "invalid memory address or nil pointer dereference" is occurring

	err := json.Unmarshal(msg.Payload, &employee)
	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %v", msg.Payload)
		return errors.Wrap(err, "failed to unmarshal json from staffing.Employee")
	}

	h.logger.Info("received event", watermill.LogFields{
		"employee_id":   employee.ID,
		"employee_name": employee.Name,
	})

	// err = h.service.CreateUser(context.Background(), cmd)
	// if err != nil {
	// 	return errors.Wrap(err, "failed to create user")
	// }
	//
	return nil
}
