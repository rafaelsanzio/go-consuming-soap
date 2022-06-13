package main

import (
	"log"
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/api"
)

func main() {
	log.Println("starting up...")
	log.Fatal(http.ListenAndServe(":8000", api.NewRouter()))
}
