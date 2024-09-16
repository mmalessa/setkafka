package kfk

/**
https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go@v1.9.2/kafka
*/

import (
	"context"
	"fmt"
	"setkafka/pkg/app"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

	if topicNameFrom == "" || topicNameTo == "" || topicNameFrom == topicNameTo {
		return fmt.Errorf("invalid topics FROM: %s, TO: %s", topicNameFrom, topicNameTo)
	}

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

	if err := k.resetTopicOffset(cc, topicNameFrom); err != nil {
		return err
	}

	// Get Topics Metadata
	metadata, err := pc.GetMetadata(&topicNameTo, false, 1000)
	if err != nil {
		return err
	}

	// Check topicTo exists
	fmt.Printf("Check topic %s exists and is empty\n", topicNameTo)
	numPartitionsTo := 0
	topicToMetadata, ok := metadata.Topics[topicNameTo]
	if ok {
		numPartitionsTo = len(topicToMetadata.Partitions)
	}
	if numPartitionsTo == 0 {
		return fmt.Errorf("topic %s not found", topicNameTo)
	}

	// Check topicTo is empty
	for _, partition := range topicToMetadata.Partitions {
		lo, hi, err := pc.QueryWatermarkOffsets(topicNameTo, 0, 100)
		if err != nil {
			return err
		}
		// fmt.Printf("P: %d, L:%d, H:%d\n", partition.ID, lo, hi)
		if msgCount := hi - lo; msgCount != 0 {
			return fmt.Errorf("topic %s exists but is not empty (partition: %d)", topicNameTo, partition.ID)
		}
	}

	fmt.Println("Copy")
	// COPY
	messageCount := 0
	deliveryChan := make(chan kafka.Event)
	for {
		msg, err := cc.ReadMessage(1000 * time.Millisecond) // TODO customize timeout -> from config
		if err != nil {
			if kafkaErr, ok := err.(kafka.Error); ok {
				if kafkaErr.Code() == kafka.ErrTimedOut {
					break
				}
			}
			return err

		}
		// fmt.Printf("Message: %s\n", string(msg.Value))
		err = pc.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topicNameTo,
			},
			Value:   msg.Value,
			Headers: msg.Headers,
			Key:     msg.Key,
		},
			deliveryChan,
		)
		if err != nil {
			return err
		}
		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return fmt.Errorf("%s", m.TopicPartition.Error.Error())
		}
		messageCount++
	}
	fmt.Printf("Total messages counted: %d\n", messageCount)
	return nil
}

func (k *Kfk) GetTopicContent(topicName string) ([]*kafka.Message, error) {
	cc, err := k.consumerConnect()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	if err := k.resetTopicOffset(cc, topicName); err != nil {
		return nil, err
	}

	messageCount := 0
	var messages []*kafka.Message
	for {
		msg, err := cc.ReadMessage(1000 * time.Millisecond) // TODO customize timeout -> from config
		if err != nil {
			if kafkaErr, ok := err.(kafka.Error); ok {
				if kafkaErr.Code() == kafka.ErrTimedOut {
					break
				}
			}
			return nil, err

		}
		messages = append(messages, msg)
		messageCount++
	}

	return messages, nil
}

// TODO - common interface Producer|Consumer?
func (k *Kfk) resetTopicOffset(consumer *kafka.Consumer, topicName string) error {

	metadata, err := consumer.GetMetadata(&topicName, false, 1000)
	if err != nil {
		return err
	}

	// Reset consumer group offset for topic
	fmt.Printf("Reset topic %s offset\n", topicName)
	numPartitionsFrom := 0
	topicMetadata, ok := metadata.Topics[topicName]
	if ok {
		numPartitionsFrom = len(topicMetadata.Partitions)
	}
	if numPartitionsFrom == 0 {
		return fmt.Errorf("topic %s not found", topicName)
	}
	var partitionsToAssign []kafka.TopicPartition
	for _, partition := range topicMetadata.Partitions {
		partitionsToAssign = append(partitionsToAssign, kafka.TopicPartition{
			Topic:     &topicName,
			Partition: partition.ID,
			Offset:    kafka.OffsetBeginning,
		})
	}
	consumer.Assign(partitionsToAssign)

	return nil
}
