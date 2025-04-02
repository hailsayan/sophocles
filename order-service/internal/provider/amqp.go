package provider

import (
	"github.com/hailsayan/sophocles/order-service/internal/config"
	"github.com/hailsayan/sophocles/order-service/internal/mq"
	pmq "github.com/hailsayan/sophocles/pkg/mq"
)

func BootstrapAMQP(cfg *config.Config) []pmq.AMQPConsumer {
	cancelNotificationProducer := mq.NewCancelNotificationProducer(rabbitmq)

	return []pmq.AMQPConsumer{
		mq.NewAutoCancelConsumer(rabbitmq, dataStore, cancelNotificationProducer),
	}
}
