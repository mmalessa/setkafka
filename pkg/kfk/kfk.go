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

func (k *Kfk) adminConnect() (*kafka.AdminClient, error) {
	adminConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.cfg.BootstrapServers,
	}

	adminClient, err := kafka.NewAdminClient(adminConfig)
	if err != nil {
		return nil, err
	}
	return adminClient, nil
}

func (k *Kfk) producerConnect() (*kafka.Producer, error) {
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.cfg.BootstrapServers,
	}

	producerClient, err := kafka.NewProducer(producerConfig)
	if err != nil {
		return nil, err
	}
	return producerClient, nil
}

func (k *Kfk) GetTopicList() (*kafka.Metadata, error) {
	pc, err := k.producerConnect()
	if err != nil {
		return nil, err
	}
	defer pc.Close()

	// Pobieranie metadanych (w tym listy tematów)
	metadata, err := pc.GetMetadata(nil, true, 10000)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (k *Kfk) CreateTopic(topicConfig kafka.TopicSpecification) error {
	ac, err := k.adminConnect()
	if err != nil {
		return err
	}
	defer ac.Close()

	results, err := ac.CreateTopics(
		context.TODO(), //FIXME
		[]kafka.TopicSpecification{topicConfig},
	)
	if err != nil {
		return err
	}

	// TODO
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to create topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s created successfully\n", result.Topic)
		}
	}
	return nil
}

func (k *Kfk) DeleteTopic(topicName string) error {
	ac, err := k.adminConnect()
	if err != nil {
		return err
	}
	defer ac.Close()

	results, err := ac.DeleteTopics(
		context.TODO(),
		[]string{topicName},
	)
	if err != nil {
		return err
	}

	// TODO
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to delete topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s deleted successfully\n", result.Topic)
		}
	}

	return nil
}
