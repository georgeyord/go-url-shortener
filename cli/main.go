package main

import (
	"github.com/georgeyord/go-url-shortener/cli/commands"
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/spf13/viper"
)

func main() {
	config.Init()

	// Provide db to commands
	db := config.InitDb()
	viper.Set("db", db)

	// Provide kafka stats writer to commands
	kafkaWriters := config.InitKafkaWriters()
	statsTopic := viper.GetString("kafka.writers.stats.topic")
	if kafkaWriters[statsTopic] == nil {
		viper.Set(statsTopic, kafkaWriters[statsTopic])
	}

	commands.Execute()
}
