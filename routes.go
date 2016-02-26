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
	Route{
		"DomainCreate",
		"POST",
		"/api/v1/domain/create",
		DomainCreate,
	},
	Route{
		"DomainList",
		"GET",
		"/api/v1/domain/user/{user_id}",
		DomainList,
	},
}
