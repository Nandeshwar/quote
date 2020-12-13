package main

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"

	"quote/pkg/binaryinfo"
	"quote/pkg/kafka"
)

func main() {
	SendMessageToKafka()
}

func SendMessageToKafka() {
	producer := kafka.NewProducer([]string{"localhost:9092"})
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	go func() {
		procudeMsg(producer)
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				procudeMsg(producer)

			}
		}
	}()

	time.Sleep(1 * time.Hour)
	ticker.Stop()
	done <- true
}

func procudeMsg(producer sarama.AsyncProducer) {
	buf := new(bytes.Buffer)
	info := newInfo()
	toTLVBytes(buf, binaryinfo.INFO_TYPE, binaryinfo.INFO_LEN, &info)
	kafka.ProduceBinaryMessage(producer, buf.Bytes())
	logrus.Info("message sent to kafka successfully")
}

func newInfo() binaryinfo.InfoMohan {

	var name [10]byte
	copy(name[:], "Mohan")

	now := time.Now()
	var createdAt [25]byte
	copy(createdAt[:], now.Format(time.RFC3339))

	return binaryinfo.InfoMohan{
		Id:        100,
		Name:      name,
		Fee:       120.20,
		CreatedAt: createdAt,
	}
}

func toTLVBytes(buf *bytes.Buffer, tlvType binaryinfo.HeaderType, tlvLength binaryinfo.HeaderLength, data interface{}) {
	// TLV type
	_ = buf.WriteByte(uint8(tlvType))

	// TLV length
	_ = buf.WriteByte(uint8(tlvLength))

	// Data
	_ = binary.Write(buf, binary.BigEndian, data)
}
