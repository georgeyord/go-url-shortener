package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/georgeyord/go-url-shortener/pkg/kafka/readers"
	"github.com/spf13/viper"
)

const role = "worker"

func main() {
	config.Init(role)
	config.PrintIntro()

	topic := viper.GetString("kafka.readers.stats.topic")
	groupId := viper.GetString("kafka.readers.stats.groupId")
	readers.RunStatsReader(topic, groupId)
}
