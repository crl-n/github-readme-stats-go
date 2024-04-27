package main

import (
	"fmt"
	"io"
	"net/http"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct {
	username string
}

func (ghClient GithubClient) GetUserData() {
	resp, err := http.Get(GithubAPIBaseURL + "/users/" + ghClient.username)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	fmt.Println(string(body))
}
