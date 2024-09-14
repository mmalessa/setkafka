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

func (k *Kfk) consumerConnect() (*kafka.Consumer, error) {

	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers":  k.cfg.BootstrapServers,
		"group.id":           k.cfg.ConsumerGroupId,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": true,
	}
	consumerClient, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		return nil, err
	}
	return consumerClient, nil
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

func (k *Kfk) CopyTopic(topicNameFrom string, topicNameTo string) error {
	cc, err := k.consumerConnect()
	if err != nil {
		return err
	}
	defer cc.Close()
	pc, err := k.producerConnect()
	if err != nil {
		return err
	}
	defer pc.Close()

	metadata, err := cc.GetMetadata(&topicNameFrom, false, 1000)
	if err != nil {
		return err
	}
	numPartitions := 0
	if topicMetadata, ok := metadata.Topics[topicNameFrom]; ok {
		numPartitions = len(topicMetadata.Partitions)
	}
	if numPartitions == 0 {
		return fmt.Errorf("topic %s not found in metadata", topicNameFrom)
	}
	fmt.Printf("Number of partitions: %d\n", numPartitions)

	if err := cc.Subscribe(topicNameFrom, nil); err != nil {
		return err
	}

	for p := 0; p < numPartitions; p++ {
		cc.Seek(kafka.TopicPartition{
			Topic:     &topicNameFrom,
			Partition: int32(p),
			Offset:    kafka.OffsetBeginning,
		}, -1)
	}

	fmt.Println("TODO TODO TODO")
	// for {
	// 	msg, err := cc.ReadMessage(-1)
	// 	if err != nil {
	// 		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	// 		continue
	// 	}
	// 	fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	// }

	return nil
}
