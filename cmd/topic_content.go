package cmd

import (
	"fmt"
	"os"
	"setkafka/pkg/app"
	"setkafka/pkg/kfk"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicContentCmd)
	topicContentCmd.Flags().StringP("name", "n", "", "Topic name")
}

var topicContentCmd = &cobra.Command{
	Use:   "content",
	Short: "Display topic content",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			logrus.Error("You must provide a topic name")
			os.Exit(0)
		}

		logrus.Info("Topic content")
		kf := kfk.NewKfk(&app.Cfg.Kafka)
		messages, err := kf.GetTopicContent(name)
		if err != nil {
			logrus.Error(err.Error())
			os.Exit(0)
		}
		for _, msg := range messages {
			fmt.Printf("%-5d [%s] %s\n-------------------------------\n", msg.TopicPartition.Offset, string(msg.Key), string(msg.Value))
		}
	},
}
