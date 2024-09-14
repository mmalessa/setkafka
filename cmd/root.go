package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"setkafka/pkg/app"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "setkafka",
		Short: "Set Kafka",
	}

	cfgFile     string
	appConfig   app.AppConfig
	kafkaConfig app.KafkaConfig
)

func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config.yaml)")
	cobra.OnInitialize(initConfig, initLogs)

	return rootCmd.Execute()
}

func initLogs() {
	logTimestampFormat := "2006-01-02 15:04:05" // https://golang.org/src/time/format.go
	switch appConfig.LogFormat {
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
	switch appConfig.LogLevel {
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
	logrus.Debugf("LogLevel: %s", appConfig.LogLevel)
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "./config.yaml"
	}
	absolutePath, err := filepath.Abs(cfgFile)
	fmt.Printf("Config file used: %s\n", absolutePath)

	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		panic(err)
	}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)
	if err := config.LoadFiles(absolutePath); err != nil {
		panic(err)
	}

	kafkaConfig = app.KafkaConfig{}
	if err := config.BindStruct("kafka", &kafkaConfig); err != nil {
		panic(err)
	}
	appConfig = app.AppConfig{}
	if err := config.BindStruct("app", &appConfig); err != nil {
		panic(err)
	}
}
