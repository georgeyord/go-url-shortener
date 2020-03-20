package main

import (
	"github.com/georgeyord/go-url-shortener/api/actions"
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func runRouter() {
	router := gin.New()

	if config.IsEnv(config.PRODUCTION) {
		log.Info().Str("mode", gin.ReleaseMode).Msg("Running Gin in production")
		gin.SetMode(gin.ReleaseMode)
	}

	db := config.InitDb()
	kafkaWriters := config.InitKafkaWriters()
	defer config.CloseKafkaWriters(kafkaWriters)
	actions.SetupMiddlewares(router, db, kafkaWriters)

	actions.MapRoutes(router)

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal().Err(err).Msg("Router failed")
	}
}
