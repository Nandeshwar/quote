package rabbitmq

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn        *Connection
	done        chan error
	routingKeys []string
	tag         string
}

func NewConsumer(
	conn *Connection,
	ctag string,
	routingKeys []string) (*Consumer, error) {

	if conn == nil {
		return nil, fmt.Errorf("rabbitmq connection is nil")
	}

	if len(routingKeys) == 0 {
		return nil, fmt.Errorf("empty routing keys")
	}

	for _, routingKey := range routingKeys {
		logrus.Infof("Binding queue %s to exchange %s with routing key %s", conn.ConsumerQueue.Name, conn.ExchangeName, routingKey)
		if err := conn.ch.QueueBind(
			conn.ConsumerQueue.Name, // name of the queue
			routingKey,              // bindingKey
			conn.ExchangeName,       // sourceExchange
			false,                   // noWait
			nil,                     // arguments
		); err != nil {
			return nil, fmt.Errorf("error=%s while binding queue=%s to exchange=%s with routing key=%s", err.Error(), conn.ConsumerQueue.Name, conn.ExchangeName, routingKey)
		}

	}

	//return consumer so we can do our own shutdown
	return &Consumer{
		conn:        conn,
		routingKeys: routingKeys,
		tag:         ctag,
		done:        make(chan error),
	}, nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	logrus.WithFields(logrus.Fields{
		"exchangeName": c.conn.ExchangeName,
		"routingKeys":  c.routingKeys,
	}).Infof("starting Consuming (consumer tag %q) message from rabbitmq", c.tag)
	deliveries, err := c.conn.ch.Consume(
		c.conn.ConsumerQueue.Name, // name
		c.tag,                     // consumerTag,
		false,                     // noAck
		false,                     // exclusive
		false,                     // noLocal
		false,                     // noWait
		nil,                       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("error while consuming rabbitmq message=%s, exhangeName=%s, routingKeys=%s", err.Error(), c.conn.ExchangeName, c.routingKeys)
	}

	return deliveries, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	//sanity checks
	if c.conn.ch == nil || c.conn == nil || c.done == nil || c.tag == "" {
		return fmt.Errorf("got an unexpected 0 value: %+v\n", c)
	}

	if err := c.conn.ch.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	defer logrus.Infof("rabbitmq shutdown success")

	// wait for done signal
	return <-c.done
}
