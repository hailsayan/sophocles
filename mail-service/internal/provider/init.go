package provider

import (
	"github.com/hailsayan/sophocles/mail-service/internal/config"
	"github.com/hailsayan/sophocles/pkg/database"
	"github.com/hailsayan/sophocles/pkg/utils/smtputils"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitmq *amqp.Connection
	mailer   smtputils.Mailer
)

func BootstrapGlobal(cfg *config.Config) {
	rabbitmq = database.NewAMQP((*database.AmqpOptions)(cfg.Amqp))
	mailer = smtputils.NewMailer()
}
