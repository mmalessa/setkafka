package cmd

import (
	"fmt"
	"setkafka/pkg/app"
	"setkafka/pkg/kfk"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicListCmd)
}

var topicListCmd = &cobra.Command{
	Use:   "list",
	Short: "Topic list",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("List topics")
		kf := kfk.NewKfk(&app.Cfg.Kafka)
		tl, err := kf.GetTopicList()
		if err != nil {
			panic(err)
		}
		fmt.Println("Topics in Kafka cluster:")
		for _, topic := range tl.Topics {
			fmt.Printf("%-30s %d\n", topic.Topic, len(topic.Partitions))
		}
	},
}
