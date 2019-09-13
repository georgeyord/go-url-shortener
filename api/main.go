package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap()

	router := gin.Default()
	setupRoutes(router)

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}
