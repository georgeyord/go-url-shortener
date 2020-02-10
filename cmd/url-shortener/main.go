package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/cmd"
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
