package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Ping command")
		fmt.Printf("%#v\n", appConfig)
	},
}
