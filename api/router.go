package main

import (
	"log"

	"github.com/georgeyord/go-url-shortener/api/actions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func runRouter(db *gorm.DB) {
	var router = gin.Default()
	actions.SetupMiddlewares(router, db)
	actions.MapRoutes(router)

	if viper.GetString("env") == "production" {
		log.Print("Running Gin in 'release' mode in production.")
		gin.SetMode(gin.ReleaseMode)
	}

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}
