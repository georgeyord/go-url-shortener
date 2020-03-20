package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	kafkalib "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func SetupMiddlewares(router *gin.Engine, db *gorm.DB, kafkaWriters map[string]*kafkalib.Writer) {
	// Write the logs to gin.DefaultWriter
	router.Use(gin.Logger())

	// Recover from any panics and writes a 500 if there was one
	router.Use(gin.Recovery())

	// Provide db to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Provide kafka stats writer to controllers
	statsTopic := viper.GetString("kafka.writers.stats.topic")
	if kafkaWriters[statsTopic] != nil {
		router.Use(func(c *gin.Context) {
			c.Set(statsTopic, kafkaWriters[statsTopic])
			c.Next()
		})
	}
}

func MapRoutes(router *gin.Engine) {
	router.GET("/", GetHelloWorld)
	router.GET("/api/list", ListUrlPairs)
	router.GET("/r/:short", Redirect)

	router.POST("/api/create", CreateUrlPair)
}
