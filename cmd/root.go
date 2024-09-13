package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "setkafka",
		Short: "Set Kafka",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
