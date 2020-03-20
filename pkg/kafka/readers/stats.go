package readers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/georgeyord/go-url-shortener/pkg/kafka"
	"github.com/rs/zerolog/log"
	kafkalib "github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
)

const retriesThreshold = 10
const sleepOnError = 2

func RunStatsReader(topic, groupId string) {
	var reader *kafkalib.Reader

	for usedRetries := 0; usedRetries < retriesThreshold; usedRetries++ {
		reader = kafka.NewReader(topic, groupId)

		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Error().Str("topic", topic).Err(err).Msg("Error while receiving message")
				break
			}

			parse(msg, topic)
		}
		log.Warn().Int("duration", sleepOnError).Int("retry", usedRetries).Msg("Sleeping...")

		reader.Close()
		time.Sleep(sleepOnError * time.Second)
	}

	defer reader.Close()
}

func parse(msg kafkalib.Message, topic string) {
	jsonValue := msg.Value
	var value map[string]string
	errJson := json.Unmarshal(jsonValue, &value)
	if errJson != nil {
		log.Error().
			Str("v", string(jsonValue)).
			Str("topic", topic).
			Err(errJson).
			Msg("Error while decoding json value")
		return
	}

	log.Info().
		Str("v", string(jsonValue)).
		Str("short", value["short"]).
		Str("long", value["long"]).
		Str("host", value["host"]).
		Time("time", msg.Time).
		Str("topic", msg.Topic).
		Int("partition", msg.Partition).
		Int64("offset", msg.Offset).
		Msg("Message received")
}
