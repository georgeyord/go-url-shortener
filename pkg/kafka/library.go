package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
	kafkalib "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"github.com/spf13/viper"
)

func NewReader(topic, groupID string) *kafkalib.Reader {
	brokers := viper.GetStringSlice("kafka.brokers")
	minBytes := viper.GetInt("kafka.bytes.min")
	maxBytes := viper.GetInt("kafka.bytes.max")

	reader := kafkalib.NewReader(kafkalib.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: minBytes,
		MaxBytes: maxBytes,
	})

	log.Printf("New kafka reader at '%v' with groupId '%s' started for topic '%s'", brokers, groupID, topic)
	return reader
}

func NewWriter(topic string) *kafkalib.Writer {
	brokers := viper.GetStringSlice("kafka.brokers")

	writer := kafkalib.NewWriter(kafkalib.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
	})

	// defer writer.Close()

	log.Printf("New kafka writer at '%v' started for topic '%s'", brokers, topic)
	return writer
}
