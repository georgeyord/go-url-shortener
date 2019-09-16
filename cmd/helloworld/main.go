package main

import (
	"github.com/georgeyord/go-scrumpoker-api/pkg/cmd"
	"github.com/georgeyord/go-scrumpoker-api/pkg/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
