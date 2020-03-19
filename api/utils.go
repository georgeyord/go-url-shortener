package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func getServerAddress() (serverAddress string) {
	if viper.IsSet("server.address") {
		serverAddress = viper.GetString("server.address")
	}
	serverAddress += fmt.Sprintf(":%s", viper.GetString("server.port"))
	return
}
