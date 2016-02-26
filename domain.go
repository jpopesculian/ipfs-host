package main

import (
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
	resp, err := DoDomainCreate(reqForm)
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
	resp, err := DoDomainList(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header.Set(AccessTokenHeader, accessToken)
	ForwardResponse(resp, w)
}

func DomainDelete(w http.ResponseWriter, r *http.Request) {
	user, accessToken, err := RequestIsAuthenticated(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	name := mux.Vars(r)["name"]
	resp, err := DoDomainOwner(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var owner Owner
	if err := ReadJsonStruct(resp.Body, &owner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Id != owner.UserId {
		http.Error(w, "You cannot access this person's domains!", http.StatusUnauthorized)
		return
	}
	resp, err = DoDomainDelete(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header.Set(AccessTokenHeader, accessToken)
	ForwardResponse(resp, w)
}
