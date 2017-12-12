package router

import (
	"github.com/MatthiasHertel/ws17_pp/src/handler"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// NewRouter init mux router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	ch := alice.New(handler.LoggingHandler)

	// ch = ch.Append(BasicAuth)

	router.Handle("/jobs", ch.ThenFunc(handler.AllJobsEndPoint)).Methods("GET")
	router.Handle("/jobs", ch.ThenFunc(handler.CreateJobEndPoint)).Methods("POST")
	router.Handle("/jobs", ch.ThenFunc(handler.UpdateJobEndPoint)).Methods("PUT")
	router.Handle("/jobs", ch.ThenFunc(handler.DeleteJobEndPoint)).Methods("DELETE")
	router.Handle("/jobs/{jobID}", ch.ThenFunc(handler.FindJobEndpoint)).Methods("GET")

	router.Handle("/", ch.ThenFunc(handler.IndexHandler))
	return router
}
