package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/georgeyord/go-scrumpoker-api/pkg/config"
	"github.com/spf13/viper"
)

func bootstrap() {
	config.Init()
	printIntro()
}

func printIntro() {
	appFigure := figure.NewFigure(viper.GetString("application.name"), viper.GetString("application.asciiart.theme"), true)
	appFigure.Print()
}
