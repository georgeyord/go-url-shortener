package main

import (
	"log"

	"github.com/georgeyord/go-url-shortener/api/actions"
	"github.com/gin-gonic/gin"
)

func serve() {
	router := getRouter()

	if err := router.Run(getServerAddress()); err != nil {
		log.Fatal(err)
	}
}

func getRouter() *gin.Engine {
	var router = gin.Default()
	mapRoutes(router)
	return router
}

func mapRoutes(router *gin.Engine) {
	router.GET("/", actions.GetHelloWorld())
}
