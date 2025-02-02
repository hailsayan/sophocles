package database

import "github.com/IBM/sarama"

type KafkaProducerOptions struct {
	SaramaConfig   *sarama.Config
	Brokers        []string
	Topic          string
	RetryMax       int
	FlushFrequency int
	Acks           sarama.RequiredAcks
	ReturnSuccess  bool
}

type KafkaConsumerOptions struct {
	SaramaConfig  *sarama.Config
	Brokers       []string
	ConsumerGroup string
	InitialOffset int64
}