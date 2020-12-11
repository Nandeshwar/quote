package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"

	"quote/pkg/binaryinfo"
	"quote/pkg/kafka"
)

func main() {
	//kafka.ConsumeMessage()
	// The code above works too. Code below is for kafka group consumer

	kafkaConsumer := kafka.NewConsumer([]string{"localhost:9092"}, []string{"quote-topic"}, "kafka-group-quote", nil)

	kafkaConsumer.ConsumeKafkaMessage(parseBinaryMsg)

}

func parseBinaryMsg(msgBytes []byte) {
	tlvHeader := &binaryinfo.Header{}

	reader := bytes.NewReader(msgBytes)

	for {
		if err := binary.Read(reader, binary.BigEndian, tlvHeader); err != nil {
			if err == io.EOF {
				logrus.Errorf("binary data parsing reached to end of file=%v", err)
				break
			}
			fmt.Printf("binary header read failed: %v", err)
			break
		}

		extractPaylaod(tlvHeader, reader)

	}
}

func extractPaylaod(tlvHeader *binaryinfo.Header, reader io.Reader) {
	var payloader binaryinfo.ByteConverter

	switch tlvHeader.Type {
	case binaryinfo.Mohan:
		payloader = &binaryinfo.InfoMohan{}

	default:
		logrus.Errorf("found unknown type while decoding bytes")
	}

	err := binary.Read(reader, binary.BigEndian, payloader)
	if err != nil {
		logrus.Errorf("error reading payload=%v", err)
	}

	jsonStr, err := payloader.ToJSON()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("Decoded string=", jsonStr)
}
