package main

import (
	// "blabla"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	// blabla.Blub("package test")

	log.Fatal(http.ListenAndServe(":8080", router))
}
