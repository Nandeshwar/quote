package main

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"quote/pkg/binaryinfo"
	"quote/pkg/kafka"
	"time"
)

func main() {
	SendMessageToKafka()
}

func SendMessageToKafka() {
	producer := kafka.NewProducer([]string{"localhost:9092"})
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				buf := new(bytes.Buffer)
				info := newInfo()
				toTLVBytes(buf, binaryinfo.INFO_TYPE, binaryinfo.INFO_LEN, info)
				kafka.ProduceBinaryMessage(producer, buf.Bytes())
				logrus.Info("message sent to kafka successfully")

			}
		}
	}()

	time.Sleep(1 * time.Hour)
	ticker.Stop()
	done <- true
}

func newInfo() binaryinfo.Info {
	now := time.Now()
	return binaryinfo.Info{
		Id:        1,
		Name:      "Mohan",
		Fee:       120.20,
		CreatedAt: now.Format(time.RFC3339),
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
