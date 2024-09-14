package app

type KafkaConfig struct {
	Dsn              string `mapstructure:"dsn"`
	BootstrapServers string `mapstructure:"bootstrap_servers"`
	SecurityProtocol string `mapstructure:"security_protocol"`
}

type AppConfig struct {
	LogFormat string `mapstructure:"log_format"`
	LogLevel  string `mapstructure:"log_level"`
}
