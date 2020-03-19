package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
	"github.com/georgeyord/go-url-shortener/pkg/kafka/readers"
)

func main() {
	config.Init()
	config.PrintIntro()

	readers.RunStatsTopic()
}
