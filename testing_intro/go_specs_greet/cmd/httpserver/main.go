package main

import (
	"go_specs_greet/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	mux := httpserver.GetServerMux()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
