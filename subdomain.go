package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleSubdomain(w http.ResponseWriter, r *http.Request) {
	var hash string
	subdomain := mux.Vars(r)["subdomain"]
	if len(subdomain) < 1 {
		hash = config.RootHash
	} else {
		resp, err := DoDomainGet(subdomain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.StatusCode != 200 {
			http.Error(w, "Could not find subdomain!", http.StatusNotFound)
			return
		}
		var domain Domain
		if err := ReadJsonStruct(resp.Body, &domain); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hash = domain.Hash
	}
	ServeFromIpfs(hash, r.RequestURI, w)
}
