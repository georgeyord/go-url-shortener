package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	bootstrap()

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
