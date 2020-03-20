package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

const Role = "api"

func main() {
	config.Init()
	config.PrintIntro(Role)

	// Application blocking, always run last
	runRouter()
}
