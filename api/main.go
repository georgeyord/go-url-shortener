package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

func main() {
	config.Init()
	printIntro()
	serve()
}
