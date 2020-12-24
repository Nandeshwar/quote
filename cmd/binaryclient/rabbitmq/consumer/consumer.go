package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"quote/pkg/rabbitmq"
)

func main() {
	//rabbitmqConsumer, err := rabbitmq.NewConsumer(rabbitmqConnection, rabbitMQTestExchange, rabbitMQConsumerTag, routingKeys)

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

	consumer, err := rabbitmq.NewConsumer(rabbitMQConnection, "nks-consumer-tag", []string{"nks-routing-key"})
	if err != nil {
		fmt.Println("error connecting consumer=", consumer)
	}

	defer func() {
		err := consumer.Shutdown()
		if err != nil {
			logrus.Infof("Error closing rabbit receiver, error: %s", err.Error())
		}
	}()

	rabbitMQCh, err := consumer.Consume()
	if err != nil {
		fmt.Println("error consuming message=", err)
	}

	for consumerMsg := range rabbitMQCh {
		fmt.Println("Message Consumed from rabbitMQ=", string(consumerMsg.Body))
		// consumerMsg.Reject(false)
		// Reject(false) -- drop message,
		// Reject(true) - queue this message to be delivered to a consumer on a
		//different channel

		// release message
		err := consumerMsg.Ack(false)
		if err != nil {
			fmt.Println("error in Ack for consumer")
		}
	}
}
