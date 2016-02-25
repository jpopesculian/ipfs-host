package main

type Domain struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
}

type Domains []Domain

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type Users []User

type AccessToken struct {
	Value string `json:"access_token"`
}

type AuthenticatedUser struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"jwt"`
}
