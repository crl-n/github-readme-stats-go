package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct {
	username string
}

func (ghClient GithubClient) makeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (ghClient GithubClient) GetUserData() {
	body, err := ghClient.makeRequest(GithubAPIBaseURL + "/users/" + ghClient.username)
	checkError(err)

	fmt.Println(string(body))
}

type Repo struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
}

func (ghClient GithubClient) GetUserRepos() {
	body, err := ghClient.makeRequest(GithubAPIBaseURL + "/users/" + ghClient.username + "/repos")
	checkError(err)

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	checkError(err)

	for _, repo := range repos {
		fmt.Println(repo)
	}
}
