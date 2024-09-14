package kfk

/**
https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go@v1.9.2/kafka
*/

import (
	"context"
	"fmt"
	"setkafka/pkg/app"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kfk struct {
	cfg *app.KafkaConfig
}

func NewKfk(kCfg *app.KafkaConfig) *Kfk {
	kfk := &Kfk{
		cfg: kCfg,
	}
	return kfk
}

func (k *Kfk) adminConnect() *kafka.AdminClient {
	adminConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.cfg.BootstrapServers,
	}

	adminClient, err := kafka.NewAdminClient(adminConfig)
	if err != nil {
		panic(err)
	}
	return adminClient
}

func (k *Kfk) producerConnect() *kafka.Producer {
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.cfg.BootstrapServers,
	}

	producerClient, err := kafka.NewProducer(producerConfig)
	if err != nil {
		panic(err)
	}
	return producerClient
}

func (k *Kfk) GetTopicList() *kafka.Metadata {
	pc := k.producerConnect()
	defer pc.Close()

	// Pobieranie metadanych (w tym listy tematów)
	metadata, err := pc.GetMetadata(nil, true, 10000)
	if err != nil {
		panic(err)
	}
	return metadata
}

func (k *Kfk) CreateTopic(topicConfig kafka.TopicSpecification) {
	ac := k.adminConnect()
	defer ac.Close()

	results, err := ac.CreateTopics(
		context.TODO(), //FIXME
		[]kafka.TopicSpecification{topicConfig},
	)
	if err != nil {
		panic(err)
	}

	// TODO
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to create topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s created successfully\n", result.Topic)
		}
	}

}

func (k *Kfk) DeleteTopic(topicName string) {
	ac := k.adminConnect()
	defer ac.Close()

	results, err := ac.DeleteTopics(
		context.TODO(),
		[]string{topicName},
	)
	if err != nil {
		panic(err)
	}

	// TODO
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to delete topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s deleted successfully\n", result.Topic)
		}
	}
}

func (k *Kfk) Debug() {

	fmt.Printf("%#v\n", k.cfg)

}
