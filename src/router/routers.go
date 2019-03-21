package router

import (
	"../handlers"
	"net/http"
)

type Route struct {
	Name        string
	//Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		//"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Trigger",
		//"GET",
		"/trigger",
		handlers.Trigger,
	},
	Route{
		"Login",
		"/login",
		handlers.UserLogin,
	},
}