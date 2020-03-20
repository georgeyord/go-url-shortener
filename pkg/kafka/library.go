package kafka

import (
	"github.com/rs/zerolog/log"
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

	log.Info().Str("topic", topic).Str("groupID", groupID).Msg("New kafka reader started")
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

	log.Info().Str("topic", topic).Msg("New kafka writer started")
	return writer
}
