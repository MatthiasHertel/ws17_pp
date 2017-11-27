package main

import (
	"net/http"
	"testing"

	. "github.com/emicklei/forest"
)

var hpc = NewClient("http://localhost:7777", new(http.Client))

func TestApiListJobs(t *testing.T) {
	get := NewConfig("/jobs", "forest").Header("Accept", "application/json")
	r := hpc.GET(t, get)
	ExpectStatus(t, r, 200)
}

func TestApiAddJobs(t *testing.T) {
	posting := NewConfig("/jobs", "forest").Header("Content-Type", "application/json").Body(`{"name" : "test"}`)
	r := hpc.POST(t, posting)
	ExpectStatus(t, r, 201)
}
