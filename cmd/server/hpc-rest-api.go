package main

import (
	"log"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp/src/router"
)

// Define HTTP request routes
func main() {
	router := router.NewRouter()

	log.Print("starting server at port 7777")

	if err := http.ListenAndServe(":7777", router); err != nil {
		log.Fatal(err)
	}
}
