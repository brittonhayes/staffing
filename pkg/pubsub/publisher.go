package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewPublisher(config kafka.PublisherConfig, logger watermill.LoggerAdapter) message.Publisher {
	publisher, err := kafka.NewPublisher(
		config,
		logger,
	)
	if err != nil {
		panic(err)
	}

	return publisher
}
