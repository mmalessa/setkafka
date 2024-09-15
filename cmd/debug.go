package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"setkafka/pkg/app"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(debugCmd)
}

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Debug")
		prettyJSON, err := json.MarshalIndent(app.Cfg, "", "  ")
		if err != nil {
			logrus.Error(err.Error())
			os.Exit(0)
		}

		fmt.Println(string(prettyJSON))

	},
}
