package readers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/georgeyord/go-url-shortener/pkg/kafka"
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
				log.Printf("Error while receiving message on topic '%s': %s", topic, err.Error())
				break
			}

			value := msg.Value

			fmt.Printf("Message at %s: %s (Topic/partition/offset: %v/%v/%v)\n", msg.Time, string(value), msg.Topic, msg.Partition, msg.Offset)
		}
		log.Printf("Sleeping for %d seconds (%d retry)", sleepOnError, usedRetries)

		reader.Close()
		time.Sleep(sleepOnError * time.Second)
	}

	defer reader.Close()
}
