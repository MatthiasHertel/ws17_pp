package main

import (
	// "blabla"
	"log"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp"
)

func main() {

	router := ws17_pp.NewRouter()

	// blabla.Blub("package test")

	log.Fatal(http.ListenAndServe(":8080", router))
}
