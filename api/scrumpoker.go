package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/georgeyord/go-scrumpoker-api/pkg/helloworld"
)

type Scrumpoker string

func (s Scrumpoker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request %v", r)
	name := "world"

	fmt.Fprintf(w, helloworld.GetHelloWorldMessage(name))
}
