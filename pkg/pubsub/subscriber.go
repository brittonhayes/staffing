package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewSubscriber(config kafka.SubscriberConfig, logger watermill.LoggerAdapter) message.Subscriber {
	sub, err := kafka.NewSubscriber(
		config,
		logger,
	)
	if err != nil {
		panic(err)
	}

	return sub
}
