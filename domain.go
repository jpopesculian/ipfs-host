package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func DomainCreate(w http.ResponseWriter, r *http.Request) {
	user, accessToken, err := RequestIsAuthenticated(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var form DomainCreateForm
	if err := ReadJsonStruct(r.Body, &form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqForm := DomainUserCreateForm{
		user.Id,
		form.Name,
		form.Hash,
	}
	reqBody, err := json.Marshal(reqForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqUrl := fmt.Sprintf("http://%s/create", config.DomainServiceHost)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header.Set(AccessTokenHeader, accessToken)
	ForwardResponse(resp, w)
}

func DomainList(w http.ResponseWriter, r *http.Request) {
	user, accessToken, err := RequestIsAuthenticated(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userId := mux.Vars(r)["user_id"]
	if user.Id != userId {
		http.Error(w, "You cannot access this person's domains!", http.StatusUnauthorized)
	}
	reqUrl := fmt.Sprintf("http://%s/user/%s", config.DomainServiceHost, userId)
	req, err := http.NewRequest("GET", reqUrl, bytes.NewReader([]byte{}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header.Set(AccessTokenHeader, accessToken)
	ForwardResponse(resp, w)
}
