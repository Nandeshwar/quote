package rabbitmq

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"net/url"
	"strings"
)

type Connection struct {
	Conn          *amqp.Connection
	ch            *amqp.Channel
	ConsumerQueue amqp.Queue
	ExchangeName  string
}

func NewConnection(
	host,
	username,
	password,
	exchangeName,
	exchangeType,
	vhost,
	queueName string) (*Connection, error) {

	if len(strings.TrimSpace(host)) == 0 {
		return nil, errors.New("no rabbitmq host provided")
	}

	_, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, fmt.Errorf("invalid rabbitmq host=%s", host)
	}

	if !strings.Contains(host, "amqp") {
		return nil, fmt.Errorf("expected protocol=amqp", host)
	}

	if len(strings.TrimSpace(exchangeName)) == 0 {
		return nil, fmt.Errorf("exchange name should not be empty")
	}

	hostArr := strings.Split(host, "//")
	url := hostArr[0] + "//" + username + ":" + password + "@" + hostArr[1]
	config := amqp.Config{
		Vhost: vhost,
	}
	conn, err := amqp.DialConfig(url, config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to rabbitmq=%s, host=%s, username=%s, exchangeName=%s, exchangeType=%s", err.Error(), host, username, exchangeName, exchangeType)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel creation error=%s while connecting to rabbitmq, host=%s, username=%s, exchangeName=%s, exchangeType=%s", err.Error(), host, username, exchangeName, exchangeType)
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("exchange declaration error=%s while connecting to rabbitmq, host=%s, username=%s, exchangeName=%s, exchangeType=%s", err.Error(), host, username, exchangeName, exchangeType)
	}

	internalQueue, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("queue declaration error=%s while connecting to rabbitmq, host=%s, username=%s, exchangeName=%s, exchangeType=%s", err.Error(), host, username, exchangeName, exchangeType)
	}

	return &Connection{
		Conn:          conn,
		ch:            ch,
		ExchangeName:  exchangeName,
		ConsumerQueue: internalQueue,
	}, nil
}

func (c *Connection) Close() {
	err := c.ch.Close()
	if err != nil {
		logrus.Fatalf("Error closing rabbitmq connection channel, error: %s", err.Error())
	}
	err = c.Conn.Close()
	if err != nil {
		logrus.Fatalf("Error closing rabbitmq client, error: %s", err.Error())
	}
}

func (c *Connection) SendMessage(routingKey string, msg []byte) error {
	if len(strings.TrimSpace(routingKey)) == 0 {
		return errors.New("no routing key provided")
	}

	err := c.ch.Publish(
		c.ExchangeName,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"exchangeName": c.ExchangeName,
			"routingKey":   routingKey,
		}).Fatal("error sending message at rabbitmq")
	}

	logrus.WithFields(logrus.Fields{
		"exchangeName": c.ExchangeName,
		"routingKey":   routingKey,
	}).Info("message sent to rabbitmq successfully")

	return nil
}
