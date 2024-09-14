package main

import (
	"github.com/sirupsen/logrus"
	"setkafka/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Error(err)
	}
}

