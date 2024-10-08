package cmd

import (
	"os"
	"setkafka/pkg/app"
	"setkafka/pkg/kfk"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicCopyCmd)
	topicCopyCmd.Flags().StringP("from", "", "", "Topic name FROM")
	topicCopyCmd.Flags().StringP("to", "", "", "Topic name TO")
}

var topicCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy data from topic to topic",
	Run: func(cmd *cobra.Command, args []string) {
		nameFrom, _ := cmd.Flags().GetString("name-from")
		nameTo, _ := cmd.Flags().GetString("name-to")
		if nameFrom == "" || nameTo == "" {
			logrus.Error("You must provide a topic names name-from and name-to")
			os.Exit(0)
		}

		logrus.Info("Copy topic")
		kf := kfk.NewKfk(&app.Cfg.Kafka)
		if err := kf.CopyTopic(nameFrom, nameTo); err != nil {
			logrus.Error(err.Error())
			os.Exit(0)
		}
	},
}
