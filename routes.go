package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Methods string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	Route{
		"UserService",
		"GET, POST",
		"/api/v1/user/{route}",
		ForwardToUserService,
	},
}
