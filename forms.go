package main

type DomainCreateForm struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type DomainUserCreateForm struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Hash   string `json:"hash"`
}
