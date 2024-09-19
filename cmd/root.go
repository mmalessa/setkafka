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
	cfgFile := "./setkafka.yaml"
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, fmt.Sprintf("config file (default is %s)", cfgFile))

	absolutePath, err := filepath.Abs(cfgFile)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		panic(err)
	}

	fmt.Printf("Config file used: %s\n", absolutePath)
	if err := app.InitConfig(absolutePath); err != nil {
		panic(err)
	}
}

func initLogs() {
	if err := app.InitLogs(); err != nil {
		panic(err)
	}
}
