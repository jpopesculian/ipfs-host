package main

import (
	"fmt"
	"net/http"
)

func ServeFromIpfs(hash, url string, w http.ResponseWriter) {
	reqUrl := fmt.Sprintf("http://%s/ipfs/%s%s", config.IpfsHost, hash, url)
	client := &http.Client{}
	resp, err := client.Get(reqUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ForwardResponse(resp, w)
}
