package ws17_pp

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"JobIndex",
		"GET",
		"/jobs",
		JobIndex,
	},
	Route{
		"JobShow",
		"GET",
		"/jobs/{jobId}",
		JobShow,
	},
	Route{
		"JobCreate",
		"POST",
		"/jobs",
		JobCreate,
	},
}
