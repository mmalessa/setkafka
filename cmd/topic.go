package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(topicCmd)
}

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Topic tools",
}
