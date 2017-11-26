// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"2_rescale_api/restapi/operations"
	"2_rescale_api/restapi/operations/jobs"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name JobList --spec ../swagger.yml

func configureFlags(api *operations.JobListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.JobListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.JobsAddOneHandler = jobs.AddOneHandlerFunc(func(params jobs.AddOneParams) middleware.Responder {
		return middleware.NotImplemented("operation jobs.AddOne has not yet been implemented")
	})
	api.JobsDestroyOneHandler = jobs.DestroyOneHandlerFunc(func(params jobs.DestroyOneParams) middleware.Responder {
		return middleware.NotImplemented("operation jobs.DestroyOne has not yet been implemented")
	})
	api.JobsFindJobsHandler = jobs.FindJobsHandlerFunc(func(params jobs.FindJobsParams) middleware.Responder {
		return middleware.NotImplemented("operation jobs.FindJobs has not yet been implemented")
	})
	api.JobsUpdateOneHandler = jobs.UpdateOneHandlerFunc(func(params jobs.UpdateOneParams) middleware.Responder {
		return middleware.NotImplemented("operation jobs.UpdateOne has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
