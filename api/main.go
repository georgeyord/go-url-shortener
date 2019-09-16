package main

import (
	"github.com/georgeyord/go-scrumpoker-api/pkg/config"
)

func main() {
	config.Init()
	printIntro()
	serve()
}
