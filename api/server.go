package main

import (
	"log"

	"github.com/georgeyord/go-url-shortener/api/actions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func serve(db *gorm.DB) {
	var router = gin.Default()
	setupMiddlewares(router, db)
	mapRoutes(router)

	if viper.GetString("env") == "production" {
		log.Print("Running Gin in 'release' mode in production.")
		gin.SetMode(gin.ReleaseMode)
	}

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}

func setupMiddlewares(router *gin.Engine, db *gorm.DB) {
	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}

func mapRoutes(router *gin.Engine) {
	router.GET("/", actions.GetHelloWorld)
	router.GET("/api/list", actions.FindUrlPairs)
	router.GET("/r/:short", actions.Redirect)

	router.POST("/api/create", actions.CreateUrlPair)
}
