package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

const role = "api"

func main() {
	config.Init(role)
	config.PrintIntro()

	// Application blocking, always run last
	runRouter()
}
