package router

import (
	. "github.com/MatthiasHertel/ws17_pp/src/handler"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	ch := alice.New(LoggingHandler)

	// ch = ch.Append(BasicAuth)

	router.Handle("/jobs", ch.ThenFunc(AllJobsEndPoint)).Methods("GET")
	router.Handle("/jobs", ch.ThenFunc(CreateJobEndPoint)).Methods("POST")
	router.Handle("/jobs", ch.ThenFunc(UpdateJobEndPoint)).Methods("PUT")
	router.Handle("/jobs", ch.ThenFunc(DeleteJobEndPoint)).Methods("DELETE")
	router.Handle("/jobs/{id}", ch.ThenFunc(FindJobEndpoint)).Methods("GET")

	router.Handle("/", ch.ThenFunc(IndexHandler))
	return router
}
