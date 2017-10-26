package main

import (
	// "blabla"
	"log"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp"
)

func main() {

	router := ws17_pp.NewRouter()

	log.Print("starting server at port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
