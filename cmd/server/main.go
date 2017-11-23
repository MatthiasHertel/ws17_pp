package main

import (
	// "blabla"
	"log"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp"
)

func main() {

	router := ws17_pp.NewRouter()

	log.Print("starting server at port 7777")

	log.Fatal(http.ListenAndServe(":7777", router))
}
