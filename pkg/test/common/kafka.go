package common

import (
	kafkalib "github.com/segmentio/kafka-go"
)

func InitTestKafkaWriters() map[string]*kafkalib.Writer {
	return map[string]*kafkalib.Writer{}
}
