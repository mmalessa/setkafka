package cmd

import (
	"setkafka/pkg/app"

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
		logrus.Debugf("%#v\n", app.Cfg)
	},
}
