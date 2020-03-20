package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/georgeyord/go-url-shortener/pkg/kafka/readers"
	"github.com/spf13/viper"
)

const Role = "worker"

func main() {
	config.Init()
	config.PrintIntro(Role)

	topic := viper.GetString("kafka.readers.stats.topic")
	groupId := viper.GetString("kafka.readers.stats.groupId")
	readers.RunStatsReader(topic, groupId)
}
