package main

import (
	"log"

	"github.com/georgeyord/go-scrumpoker-api/api/actions"
	"github.com/gin-gonic/gin"
)

func serve() {
	var router = gin.Default()
	mapRoutes(router)

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}

func mapRoutes(router *gin.Engine) {
	router.GET("/", actions.GetHelloWorld())
}
