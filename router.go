package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var router *mux.Router

func InitRouter() {
	router = mux.NewRouter().StrictSlash(true)
	var subdomainHandler http.HandlerFunc = HandleSubdomain
	router.
		Methods("GET").
		Host(fmt.Sprintf("{subdomain}.%s", config.Host)).
		Handler(Logger(subdomainHandler, "Subdomain"))

	for _, route := range routes {
		router.
			Methods(strings.Split(route.Methods, ", ")...).
			Path(route.Path).
			Name(route.Name).
			Host(config.Host).
			Handler(Logger(route.Handler, route.Name))
	}

	router.
		Methods("GET").
		Host(config.Host).
		Handler(Logger(subdomainHandler, "Subdomain"))
}
