package main

import (
	"log"
	"net/http"
)

func main() {
	const serverAddress = ":8080"
	log.Printf("Start serving on %s...", serverAddress)
	http.HandleFunc("/", scrumpoker)
	err := http.ListenAndServe(serverAddress, nil)

	if err != nil {
		log.Fatal(err)
	}
}
