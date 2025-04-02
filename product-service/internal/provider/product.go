package provider

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/sophocles/pkg/database"
	"github.com/hailsayan/sophocles/product-service/internal/config"
	"github.com/hailsayan/sophocles/product-service/internal/constant"
	"github.com/hailsayan/sophocles/product-service/internal/controller"
	"github.com/hailsayan/sophocles/product-service/internal/mq"
	"github.com/hailsayan/sophocles/product-service/internal/usecase"
)

func BootstrapProduct(cfg *config.Config, router *gin.Engine) {
	kafkaAdmin.CreateTopic(constant.ProductCreatedTopic, constant.ProductCreatedTopicNumPartitions, constant.ProductCreatedTopicReplicationFactor)
	kafkaAdmin.CreateTopic(constant.ProductDeletedTopic, constant.ProductDeletedTopicNumPartitions, constant.ProductDeletedTopicReplicationFactor)
	kafkaAdmin.CreateTopic(constant.ProductUpdatedTopic, constant.ProductUpdatedTopicNumPartitions, constant.ProductUpdatedTopicReplicationFactor)

	productCreatedProducer := mq.NewProductCreatedProducer(database.NewKafkaAsyncProducer(&database.KafkaProducerOptions{
		Brokers:        cfg.Kafka.Brokers,
		RetryMax:       cfg.Kafka.RetryMax,
		FlushFrequency: cfg.Kafka.FlushFrequency,
		ReturnSuccess:  cfg.Kafka.ReturnSuccess,
		Topic:          constant.ProductCreatedTopic,
		Acks:           sarama.WaitForLocal,
	}))
	productUpdatedProducer := mq.NewProductUpdatedProducer(database.NewKafkaSyncProducer(&database.KafkaProducerOptions{
		Brokers:        cfg.Kafka.Brokers,
		RetryMax:       cfg.Kafka.RetryMax,
		FlushFrequency: cfg.Kafka.FlushFrequency,
		ReturnSuccess:  cfg.Kafka.ReturnSuccess,
		Topic:          constant.ProductUpdatedTopic,
		Acks:           sarama.WaitForLocal,
	}))
	productDeletedProducer := mq.NewProductDeletedProducer(database.NewKafkaAsyncProducer(&database.KafkaProducerOptions{
		Brokers:        cfg.Kafka.Brokers,
		RetryMax:       cfg.Kafka.RetryMax,
		FlushFrequency: cfg.Kafka.FlushFrequency,
		ReturnSuccess:  cfg.Kafka.ReturnSuccess,
		Topic:          constant.ProductDeletedTopic,
		Acks:           sarama.WaitForLocal,
	}))
	productUseCase := usecase.NewProductUseCase(dataStore, productCreatedProducer, productUpdatedProducer, productDeletedProducer)
	productController := controller.NewProductController(productUseCase)

	productController.Route(router)
}
