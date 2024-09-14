package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"setkafka/pkg/app"

	"github.com/sirupsen/logrus"
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
	logTimestampFormat := "2006-01-02 15:04:05" // https://golang.org/src/time/format.go
	switch app.Cfg.App.LogFormat {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint:     false,
			TimestampFormat: logTimestampFormat,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: logTimestampFormat,
		})
	}
	switch app.Cfg.App.LogLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default: // error
		logrus.SetLevel(logrus.ErrorLevel)
	}
	logrus.Debugf("LogLevel: %s", app.Cfg.App.LogLevel)
}
