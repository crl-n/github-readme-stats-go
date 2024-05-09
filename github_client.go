package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/crl-n/github-readme-stats-go/logger"
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

type RepoAPIResponse struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
	PushedAt     string `json:"pushed_at"`
}

type LanguageAPIResponse map[string]int

type Repo struct {
	Name      string
	Languages map[string]int
	PushedAt  time.Time
}

func (rawRepo RepoAPIResponse) ToRepo(ghClient GithubClient) (Repo, error) {
	body, err := ghClient.makeRequest(rawRepo.LanguagesUrl)
	if err != nil {
		return Repo{}, err
	}

	var languages LanguageAPIResponse
	err = json.Unmarshal(body, &languages)
	if err != nil {
		return Repo{}, err
	}

	pushedAtTime, err := time.Parse(time.RFC3339, rawRepo.PushedAt)
	if err != nil {
		return Repo{}, err
	}

	repo := Repo{rawRepo.Name, languages, pushedAtTime}

	return repo, nil
}

func (ghClient GithubClient) GetUserRepos() ([]Repo, error) {
	body, err := ghClient.makeRequest(GithubAPIBaseURL + "/users/" + ghClient.username + "/repos")
	if err != nil {
		return nil, err
	}

	var rawRepos []RepoAPIResponse
	err = json.Unmarshal(body, &rawRepos)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	cachedRepos := RetrieveCachedRepos()

	for _, rawRepo := range rawRepos {
		rawRepoPushedAtTime, err := time.Parse(time.RFC3339, rawRepo.PushedAt)
		if err != nil {
			return nil, err
		}

		cachedRepo, found := findRepo(cachedRepos, rawRepo)

		if found && cachedRepo.PushedAt.Equal(rawRepoPushedAtTime) {
			logger.Debugf("Using cached repo data for %v\n", rawRepo.Name)
			repos = append(repos, *cachedRepo)
		} else {
			logger.Debugf("Using new repo data for %v\n", rawRepo.Name)
			repo, err := rawRepo.ToRepo(ghClient)
			if err != nil {
				return nil, err
			}
			repos = append(repos, repo)
		}
	}

	CacheRepos(repos)

	return repos, nil
}
