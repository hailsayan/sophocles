package provider

import (
	"github.com/IBM/sarama"
	"github.com/hailsayan/sophocles/pkg/database"
	pmq "github.com/hailsayan/sophocles/pkg/mq"
	"github.com/hailsayan/sophocles/product-service/internal/config"
	"github.com/hailsayan/sophocles/product-service/internal/constant"
	"github.com/hailsayan/sophocles/product-service/internal/mq"
)

func BootstrapKafka(cfg *config.KafkaConfig) []pmq.KafkaConsumer {
	return []pmq.KafkaConsumer{
		mq.NewOrderCreatedConsumer(
			database.NewKafkaConsumerGroup(&database.KafkaConsumerOptions{
				Brokers:       cfg.Brokers,
				ConsumerGroup: constant.OrderCreatedConsumerGroup,
				InitialOffset: sarama.OffsetOldest,
			}),
			dataStore,
		),
		mq.NewOrderCancelledConsumer(
			database.NewKafkaConsumerGroup(&database.KafkaConsumerOptions{
				Brokers:       cfg.Brokers,
				ConsumerGroup: constant.OrderCancelledConsumerGroup,
				InitialOffset: sarama.OffsetOldest,
			}),
			dataStore,
		),
	}
}
