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

func (ghClient GithubClient) GetUserData() {
	resp, err := http.Get(GithubAPIBaseURL + "/users/" + ghClient.username)
	checkError(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

type Repo struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
}

func (ghClient GithubClient) GetUserRepos() {
	resp, err := http.Get(GithubAPIBaseURL + "/users/" + ghClient.username + "/repos")
	checkError(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	checkError(err)

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	checkError(err)

	for _, repo := range repos {
		fmt.Println(repo)
	}
}
