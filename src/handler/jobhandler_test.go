package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func Router1() *mux.Router {
	router1 := mux.NewRouter()
	router1.HandleFunc("/jobs", AllJobsEndPoint).Methods("GET")
	return router1
}

// https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/
func TestAllJobsEndPoint(t *testing.T) {
	request1, _ := http.NewRequest("GET", "/jobs", nil)
	rr := httptest.NewRecorder()
	Router1().ServeHTTP(rr, request1)

	fmt.Print(rr.Body)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func Router2() *mux.Router {
	router2 := mux.NewRouter()
	ch := alice.New()
	// router2.Handle("/jobs/5a3184949ebea40b903a1e55", FindJobEndpoint).Methods("GET")
	router2.Handle("/jobs/{jobID}", ch.ThenFunc(FindJobEndpoint)).Methods("GET")
	return router2
}

// https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/
func TestFindJobEndpoint(t *testing.T) {
	request2, _ := http.NewRequest("GET", "/jobs/5a3184949ebea40b903a1e55", nil)
	rr2 := httptest.NewRecorder()
	Router2().ServeHTTP(rr2, request2)

	fmt.Print(rr2.Body)

	// Check the status code is what we expect.
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
