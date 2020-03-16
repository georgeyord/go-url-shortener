package main

import (
	"github.com/georgeyord/go-url-shortener/cli/commands"
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/spf13/viper"
)

func main() {
	db := config.Init()
	viper.Set("db", db)
	commands.Execute()
}
