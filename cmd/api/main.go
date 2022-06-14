package main

import (
	"log"
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/api"
)

func main() {
	log.Println("starting up on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", api.NewRouter()))
}
