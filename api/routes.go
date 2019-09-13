package main

import (
	"github.com/georgeyord/go-scrumpoker-api/pkg/helloworld"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/", GetHelloWorld())
}

func GetHelloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")

		if len(name) < 1 {
			c.JSON(400, gin.H{
				"error": "Url Param 'name' is missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": helloworld.GetHelloWorldMessage(name),
		})
	}
}
