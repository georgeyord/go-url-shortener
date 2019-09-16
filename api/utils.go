package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/viper"
)

func getServerAddress() (serverAddress string) {
	if viper.IsSet("server.address") {
		serverAddress = viper.GetString("server.address")
	}
	serverAddress += fmt.Sprintf(":%s", viper.GetString("server.port"))
	return
}

func printIntro() {
	appFigure := figure.NewFigure(viper.GetString("application.name"), viper.GetString("application.asciiart.theme"), true)
	appFigure.Print()
}
