package main

import (
	"github.com/georgeyord/go-scrumpoker-api/pkg/cli"
	"github.com/georgeyord/go-scrumpoker-api/pkg/helloworld"
	"github.com/logrusorgru/aurora"
)

func main() {
	name := cli.GetInput("name")
	message := helloworld.GetHelloWorldMessage(name)
	cli.PrintMessage(message, aurora.Magenta)
}
