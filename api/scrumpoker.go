package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/georgeyord/go-scrumpoker-api/pkg/helloworld"
)

func scrumpoker(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request %v", r)
	var name string
	nameQuery, ok := r.URL.Query()["name"]

	if !ok || len(nameQuery[0]) < 1 {
		log.Println("Url Param 'name' is missing")
	} else {
		name = string(nameQuery[0])
	}

	fmt.Fprintf(w, helloworld.GetHelloWorldMessage(name))
}
