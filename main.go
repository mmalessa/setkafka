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

// import (
// 	"net"
// 	"strconv"

// 	kafka "github.com/segmentio/kafka-go"
// )

// // https://stackoverflow.com/questions/61618623/how-to-create-kafka-topic-using-segmentios-kafka-go
// // https://github.com/segmentio/kafka-go

// func main() {

// 	conn, err := kafka.Dial("tcp", "kafka:9092")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer conn.Close()

// 	controller, err := conn.Controller()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer controllerConn.Close()

// 	topicConfigs := []kafka.TopicConfig{{Topic: "sometopic", NumPartitions: 4, ReplicationFactor: 1}}

// 	err = controllerConn.CreateTopics(topicConfigs...)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
