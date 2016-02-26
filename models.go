package main

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type Owner struct {
	UserId string `json:"user_id"`
}

type Domain struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
}
