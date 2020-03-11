package actions

import (
	"github.com/georgeyord/go-url-shortener/pkg/helloworld"
	"github.com/gin-gonic/gin"
)

func GetHelloWorld(c *gin.Context) {
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
