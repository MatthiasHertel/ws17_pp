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

	// decorator for authentication
	// ch = ch.Append(BasicAuth)

	// job resource routes
	// TODO USER PREFIX
	router.Handle("/jobs", ch.ThenFunc(handler.AllJobsEndPoint)).Methods("GET")
	router.Handle("/jobs", ch.ThenFunc(handler.CreateJobEndPoint)).Methods("POST")
	router.Handle("/jobs", ch.ThenFunc(handler.UpdateJobEndPoint)).Methods("PUT")
	router.Handle("/jobs", ch.ThenFunc(handler.DeleteJobEndPoint)).Methods("DELETE")
	router.Handle("/jobs/{jobID}", ch.ThenFunc(handler.FindJobEndpoint)).Methods("GET")
	router.Handle("/jobs/{jobID}/submit", ch.ThenFunc(handler.SubmitJobEndPoint)).Methods("GET")

	// template resource routes
	// TODO USER PREFIX
	router.Handle("/jobs/{jobID}/templates", ch.ThenFunc(handler.AllTemplatesEndPoint)).Methods("GET")
	router.Handle("/jobs/{jobID}/templates", ch.ThenFunc(handler.CreateTemplateEndPoint)).Methods("POST")
	router.Handle("/jobs/{jobID}/templates", ch.ThenFunc(handler.UpdateTemplateEndPoint)).Methods("PUT")
	router.Handle("/jobs/{jobID}/templates", ch.ThenFunc(handler.DeleteTemplateEndPoint)).Methods("DELETE")
	router.Handle("/jobs/{jobID}/templates/{templateID}", ch.ThenFunc(handler.FindTemplateEndpoint)).Methods("GET")

	// public routes
	router.Handle("/", ch.ThenFunc(handler.IndexHandler))
	return router
}
