package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

func main() {
	config.Init()
	config.PrintIntro()

	// Application blocking, always run last
	runRouter()
}
