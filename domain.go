package main

import (
	"fmt"
	"net/http"
)

func DomainCreate(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("http://%s/create", config.DomainServiceHost)
	RedirectRequestResponse(requestUrl, r, w)
}
