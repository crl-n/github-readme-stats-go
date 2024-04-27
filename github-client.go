package main

import (
	"fmt"
	"io"
	"net/http"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct{}

func (ghClient GithubClient) GetUserData(username string) {
	resp, err := http.Get(GithubAPIBaseURL + "/users/" + username)
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
