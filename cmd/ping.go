package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping",
	Long:  `Ping`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Ping command")
		// config := viper.GetStringMap("database")
		// fmt.Println(config)
	},
}
