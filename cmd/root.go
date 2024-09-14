package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"setkafka/pkg/app"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "setkafka",
		Short: "Set Kafka",
	}
)

func Execute() error {
	cobra.OnInitialize(initConfig, initLogs)
	return rootCmd.Execute()
}

func initConfig() {
	cfgFile := "./config.yaml"
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file (default is ./config.yaml)")

	absolutePath, err := filepath.Abs(cfgFile)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		panic(err)
	}

	fmt.Printf("Config file used: %s\n", absolutePath)
	app.InitConfig(absolutePath)
}

func initLogs() {
	app.InitLogs()
}
