package main

import (
	"github.com/gorilla/mux"
	"strings"
)

var router *mux.Router

func InitRouter() {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(strings.Split(route.Methods, ", ")...).
			Path(route.Path).
			Name(route.Name).
			Handler(Logger(route.Handler, route.Name))

	}
}
