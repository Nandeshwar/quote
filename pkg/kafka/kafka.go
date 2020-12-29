package kafka

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"os/signal"
	"quote/pkg/binaryinfo"
	"strings"
	"sync"
	"syscall"
	"time"
)

func NewProducer(brokerList []string) sarama.AsyncProducer {

	// For the access log, we are looking for AP semantics, with high throughput.
	// By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		logrus.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			logrus.Errorf("Failed to write access log entry=%v", err)
		}
	}()

	return producer
}

func ProduceBinaryMessage(producer sarama.AsyncProducer, message []byte) {

	producer.Input() <- &sarama.ProducerMessage{
		Topic: "quote-topic",
		Key:   sarama.StringEncoder("quote-key"),
		Value: sarama.ByteEncoder(message),
	}
}

func ConsumeMessage() {

	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:9092"}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	//topics, _ := master.Topics()

	consumer, errors := consume([]string{"quote-topic"}, master)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-consumer:
				msgCount++
				//fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				m := msg.Value
				tlvHeader := &binaryinfo.Header{}

				reader := bytes.NewReader(m)

				if err := binary.Read(reader, binary.BigEndian, tlvHeader); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("binary header read failed: %v", err)
				}
				fmt.Println("Tlv header", tlvHeader)

			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
				doneCh <- struct{}{}
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[1], sarama.OffsetOldest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		fmt.Println(" Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					fmt.Println("consumerError: ", consumerError.Err)

				case msg := <-consumer.Messages():
					consumers <- msg
					fmt.Println("Got message on topic ", topic, msg.Value)
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}

// Consumer Group logic---------------------
//https://github.com/Shopify/sarama/blob/master/examples/consumergroup/main.go
type Consumer struct {
	ready   chan bool
	Topics  []string
	Config  *sarama.Config
	Hosts   []string
	GroupID string

	ParseBinaryDataFunc func([]byte)
}

// topics: comma separated topics
func NewConsumer(hosts []string, topics []string, groupID string, config *sarama.Config) *Consumer {
	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		logrus.Panicf("Error parsing Kafka version: %v", err)
	}

	if config == nil {
		config = sarama.NewConfig()
	}

	if groupID == "" {
		groupID = "Kafka-Group-Quote"
	}

	if hosts == nil {
		hosts = []string{"localhost:9092"}
	}

	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	assignor := "range"

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	return &Consumer{
		ready:   make(chan bool),
		Topics:  topics,
		Config:  config,
		Hosts:   hosts,
		GroupID: groupID,
	}
}

func (c *Consumer) consumeMsgStarter() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(c.Hosts, c.GroupID, c.Config)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			err := client.Consume(ctx, c.Topics, c)
			if err != nil {
				logrus.Panicf("Error from consumer: %v", err)
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready // Await till the consumer has been set up
	logrus.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func (c *Consumer) ConsumeKafkaMessage(f func(msgBytes []byte)) {
	c.ParseBinaryDataFunc = f
	c.consumeMsgStarter()

}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {

		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		m := message.Value
		consumer.ParseBinaryDataFunc(m)
		session.MarkMessage(message, "")
	}

	return nil
}
