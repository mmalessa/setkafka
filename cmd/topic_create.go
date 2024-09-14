package cmd

import (
	"fmt"
	"os"
	"setkafka/pkg/app"
	"setkafka/pkg/kfk"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicCreateCmd)
	topicCreateCmd.Flags().StringP("name", "n", "", "Topic name")
	topicCreateCmd.Flags().IntP("partitions", "p", 1, "Number of partitions")
	topicCreateCmd.Flags().IntP("replications", "x", 1, "Replication factor")
	topicCreateCmd.Flags().IntP("retention", "r", 604800000, "Retention time")
}

var topicCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Topic create",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			logrus.Error("You must provide at least a topic name")
			os.Exit(0)
		}
		numPartitions, _ := cmd.Flags().GetInt("partitions")
		replicationFactor, _ := cmd.Flags().GetInt("replications")

		logrus.Info("Create topic")
		kf := kfk.NewKfk(&app.Cfg.Kafka)

		retention, _ := cmd.Flags().GetInt("retention")
		config := map[string]string{
			"retention.ms": fmt.Sprintf("%d", retention),
		}
		// replicaAssignment :=  // TODO - do we need this?

		topicConfig := kafka.TopicSpecification{
			Topic:             name,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
			// ReplicaAssignment: replicaAssignment,
			Config: config,
		}

		if err := kf.CreateTopic(topicConfig); err != nil {
			panic(err)
		}
	},
}
