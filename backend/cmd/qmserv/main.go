package main

import (
	"log"
	"net/http"

	"github.com/kahlys/quidditch/backend/api"
)

func main() {
	router := api.Handler()

	log.Fatal(http.ListenAndServe(":8000", router))
}
