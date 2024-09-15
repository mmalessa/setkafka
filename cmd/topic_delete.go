package cmd

import (
	"os"
	"setkafka/pkg/app"
	"setkafka/pkg/kfk"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicDeleteCmd)
	topicDeleteCmd.Flags().StringP("name", "n", "", "Topic name")
}

var topicDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Topic delete",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			logrus.Error("You must provide a topic name")
			os.Exit(0)
		}
		logrus.Info("Delete topic")
		kf := kfk.NewKfk(&app.Cfg.Kafka)
		if err := kf.DeleteTopic(name); err != nil {
			logrus.Error(err.Error())
			os.Exit(0)
		}
	},
}
