package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct {}

// Keys are language names, values are number of bytes of code written
type RepoLanguages map[string]int

// Processed repository enriched with language data
type Repo struct {
	Name      string
	Languages map[string]int
	PushedAt  time.Time
}

func NewGithubClient() GithubClient {
	return GithubClient{}
}

func (ghClient GithubClient) makeRequest(urlPath string) ([]byte, error) {
	resp, err := http.Get(GithubAPIBaseURL + urlPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			fmt.Sprintf(
				"Request by GithubClient to '%s' failed with status %s",
				urlPath,
				resp.Status,
			),
		)
	}

	return body, nil
}

// Fetches languages used in a repository. See:
// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repository-languages
func (ghClient GithubClient) GetRepoLanguages(username string, repo string) (RepoLanguages, error) {
	body, err := ghClient.makeRequest(
		"/repos/" + username + "/" + repo + "/languages",
	)
	if err != nil {
		return RepoLanguages{}, err
	}

	var languages RepoLanguages
	err = json.Unmarshal(body, &languages)
	if err != nil {
		return RepoLanguages{}, err
	}

	return languages, nil
}

// Fetches list of public repositories for a user.
// See: https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repositories-for-a-user
func (ghClient GithubClient) GetPublicReposList(username string) ([]RawPublicRepo, error) {
	body, err := ghClient.makeRequest("/users/" + username + "/repos")
	if err != nil {
		return nil, err
	}

	var rawPublicRepos []RawPublicRepo
	err = json.Unmarshal(body, &rawPublicRepos)
	if err != nil {
		return nil, err
	}

	return rawPublicRepos, nil
}
