package main

import (
	"log"

	"github.com/georgeyord/go-url-shortener/api/actions"
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func runRouter() {
	router := gin.Default()

	if viper.GetString("env") == "production" {
		log.Print("Running Gin in 'release' mode in production.")
		gin.SetMode(gin.ReleaseMode)
	}

	db := config.InitDb()
	kafkaWriters := config.InitKafkaWriters()
	defer config.CloseKafkaWriters(kafkaWriters)
	actions.SetupMiddlewares(router, db, kafkaWriters)

	actions.MapRoutes(router)

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}
