package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func DoDomainCreate(reqForm DomainUserCreateForm) (*http.Response, error) {
	var resp *http.Response
	reqBody, err := json.Marshal(reqForm)
	if err != nil {
		return resp, err
	}
	reqUrl := fmt.Sprintf("http://%s/create", config.DomainServiceHost)
	client := &http.Client{}
	resp, err = client.Post(reqUrl, "application/json", bytes.NewBuffer(reqBody))
	return resp, err
}

func DoDomainList(userId string) (*http.Response, error) {
	var resp *http.Response
	reqUrl := fmt.Sprintf("http://%s/user/%s", config.DomainServiceHost, userId)
	client := &http.Client{}
	resp, err := client.Get(reqUrl)
	return resp, err
}

func DoDomainOwner(name string) (*http.Response, error) {
	var resp *http.Response
	reqUrl := fmt.Sprintf("http://%s/%s/user", config.DomainServiceHost, name)
	client := &http.Client{}
	resp, err := client.Get(reqUrl)
	return resp, err
}

func DoDomainDelete(name string) (*http.Response, error) {
	var resp *http.Response
	reqUrl := fmt.Sprintf("http://%s/%s", config.DomainServiceHost, name)
	req, err := http.NewRequest("DELETE", reqUrl, bytes.NewReader([]byte{}))
	if err != nil {
		return resp, err
	}
	client := &http.Client{}
	resp, err = client.Do(req)
	return resp, err
}

func DoDomainGet(name string) (*http.Response, error) {
	var resp *http.Response
	reqUrl := fmt.Sprintf("http://%s/%s", config.DomainServiceHost, name)
	client := &http.Client{}
	resp, err := client.Get(reqUrl)
	return resp, err
}
