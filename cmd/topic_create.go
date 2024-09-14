package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	topicCmd.AddCommand(topicCreateCmd)
}

var topicCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Topic create",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Topic create")
		fmt.Printf("%#v\n", kafkaConfig)
	},
}
