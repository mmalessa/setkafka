package app

import "github.com/sirupsen/logrus"

func InitLogs(logLevel string) error {
	logTimestampFormat := "2006-01-02 15:04:05" // https://golang.org/src/time/format.go
	switch Cfg.App.LogFormat {
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

	if logLevel == "" {
		logLevel = Cfg.App.LogLevel
	}

	switch logLevel {
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
	logrus.Debugf("LogLevel: %s", Cfg.App.LogLevel)

	return nil
}
