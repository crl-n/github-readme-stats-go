package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

type RepoAPIResponse struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
	PushedAt     string `json:"pushed_at"`
}

func (rawRepo RepoAPIResponse) ToRepo(ghClient GithubClient) (Repo, error) {
	body, err := ghClient.makeRequest(rawRepo.LanguagesUrl)
	checkError(err)

	var languages LanguageAPIResponse
	err = json.Unmarshal(body, &languages)
	if err != nil {
		fmt.Println(err)
		fmt.Println(body)
		os.Exit(1)
	}

	pushedAtTime, err := time.Parse(time.RFC3339, rawRepo.PushedAt)
	checkError(err)

	repo := Repo{rawRepo.Name, languages, pushedAtTime}

	return repo, nil
}

type LanguageAPIResponse map[string]int

type Repo struct {
	Name      string
	Languages map[string]int
	PushedAt  time.Time
}

func (ghClient GithubClient) GetUserRepos() []Repo {
	body, err := ghClient.makeRequest(GithubAPIBaseURL + "/users/" + ghClient.username + "/repos")
	checkError(err)

	var rawRepos []RepoAPIResponse
	err = json.Unmarshal(body, &rawRepos)
	checkError(err)

	var repos []Repo
	for _, rawRepo := range rawRepos {
		repo, err := rawRepo.ToRepo(ghClient)
		checkError(err)
		repos = append(repos, repo)
	}

	return repos
}

func (ghClient GithubClient) GetLanguageStats() {
	repos := ghClient.GetUserRepos()

	fmt.Println(repos)
}
