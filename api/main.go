package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/viper"
)

func main() {
	bootstrap()
	appFigure := figure.NewFigure(viper.GetString("application.name"), viper.GetString("application.asciiart.theme"), true)
	appFigure.Print()

	serverAddress := getServerAddress()

	log.Printf("Start serving on %s...", serverAddress)
	http.HandleFunc("/", scrumpoker)

	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func getServerAddress() (serverAddress string) {
	if viper.IsSet("server.address") {
		serverAddress = viper.GetString("server.address")
	}
	serverAddress += fmt.Sprintf(":%s", viper.GetString("server.port"))
	return
}
