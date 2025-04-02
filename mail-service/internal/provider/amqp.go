package provider

import (
	"github.com/hailsayan/sophocles/mail-service/internal/config"
	"github.com/hailsayan/sophocles/mail-service/internal/feign"
	"github.com/hailsayan/sophocles/mail-service/internal/mq"
	pmq "github.com/hailsayan/sophocles/pkg/mq"
)

func BootstrapAMQP(cfg *config.Config) []pmq.AMQPConsumer {
	orderClient := feign.NewOrderClient(cfg.Feign.OrderURL)

	return []pmq.AMQPConsumer{
		mq.NewSendVerificationConsumer(rabbitmq, mailer),
		mq.NewAccountVerifiedConsumer(rabbitmq, mailer),
		mq.NewPaymentReminderConsumer(rabbitmq, mailer, orderClient),
		mq.NewOrderCancelConsumer(rabbitmq, mailer, orderClient),
		mq.NewOrderSuccessConsumer(rabbitmq, mailer, orderClient),
	}
}
