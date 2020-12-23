package main

import (
	"fmt"
	"quote/pkg/rabbitmq"
)

func main() {
	rabbitMQConnection, err := rabbitmq.NewConnection(
		"amqp://localhost:5672",
		"rabbitmq",
		"rabbitmq",
		"nks-exchange",
		"topic",
		"/",
		"nks-queue",
	)
	if err != nil {
		fmt.Println("error connecting to rabbitmq", err)
		return
	}

	rabbitMQConnection.SendMessage("nks-routing-key", []byte("Nandeshwar in rabbitMQ"))

}
