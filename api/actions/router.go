package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupMiddlewares(router *gin.Engine, db *gorm.DB) {
	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}

func MapRoutes(router *gin.Engine) {
	router.GET("/", GetHelloWorld)
	router.GET("/api/list", ListUrlPairs)
	router.GET("/r/:short", Redirect)

	router.POST("/api/create", CreateUrlPair)
}
