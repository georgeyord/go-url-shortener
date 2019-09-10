package main

import (
	"log"
	"net/http"
)

func main() {
	const serverAddress = "0.0.0.0:8080"
	log.Printf("Start serving on %s...", serverAddress)
	err := http.ListenAndServe(serverAddress, Scrumpoker("World"))

	if err != nil {
		log.Fatal(err)
	}
}
