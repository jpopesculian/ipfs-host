package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const AccessTokenHeader = "X-Authorization"

func RequestIsAuthenticated(r *http.Request) (User, string, error) {
	var user User
	url := fmt.Sprintf("http://%s/authenticated", config.UserServiceHost)
	req, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return user, "", err
	}
	req.Header.Set(AccessTokenHeader, r.Header.Get(AccessTokenHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return user, "", err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return user, "", err
		}
		return user, "", errors.New(string(body))
	}
	if err := ReadJsonStruct(resp.Body, &user); err != nil {
		return user, "", err
	}
	accessToken := resp.Header.Get(AccessTokenHeader)
	return user, accessToken, nil
}
