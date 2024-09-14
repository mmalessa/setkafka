package app

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

type KafkaConfig struct {
	Dsn              string `mapstructure:"dsn"`
	BootstrapServers string `mapstructure:"bootstrap_servers"`
	SecurityProtocol string `mapstructure:"security_protocol"`
}

type AppConfig struct {
	LogFormat string `mapstructure:"log_format"`
	LogLevel  string `mapstructure:"log_level"`
}

type Config struct {
	App   AppConfig
	Kafka KafkaConfig
}

var (
	Cfg Config
)

func InitConfig(cfgFile string) error {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)
	if err := config.LoadFiles(cfgFile); err != nil {
		return err
	}

	Cfg.Kafka = KafkaConfig{}
	if err := config.BindStruct("kafka", &Cfg.Kafka); err != nil {
		return err
	}
	Cfg.App = AppConfig{}
	if err := config.BindStruct("app", &Cfg.App); err != nil {
		return err
	}

	return nil
}
