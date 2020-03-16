package main

import (
	"github.com/georgeyord/go-url-shortener/pkg/config"
)

func main() {
	db := config.Init()
	printIntro()
	runRouter(db)
}
