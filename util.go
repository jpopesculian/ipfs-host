package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadJsonStruct(data io.ReadCloser, form interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(data, 1048576))
	if err != nil {
		return err
	}
	if err := data.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, form); err != nil {
		return err
	}
	return nil
}

func WriteJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
}

func IncomingToOutgoingRequest(url string, incoming *http.Request) (*http.Request, error) {
	method := incoming.Method
	if len(method) == 0 {
		method = "GET"
	}
	return http.NewRequest(method, url, io.LimitReader(incoming.Body, 1048576))
}

func RedirectRequest(url string, incoming *http.Request) (*http.Response, error) {
	var response *http.Response
	ougoing, err := IncomingToOutgoingRequest(url, incoming)
	if err != nil {
		return response, nil
	}
	ougoing.Header = incoming.Header
	client := &http.Client{}
	return client.Do(ougoing)
}

func ForwardResponse(resp *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("X-Authorization", resp.Header.Get("X-Authorization"))
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func RedirectRequestResponse(url string, r *http.Request, w http.ResponseWriter) {
	resp, err := RedirectRequest(url, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ForwardResponse(resp, w)
}
