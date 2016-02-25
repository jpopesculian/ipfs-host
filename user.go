package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ForwardToUserService(w http.ResponseWriter, r *http.Request) {
	route := mux.Vars(r)["route"]
	requestUrl := fmt.Sprintf("http://%s/%s", config.UserServiceHost, route)
	RedirectRequestResponse(requestUrl, r, w)
}
