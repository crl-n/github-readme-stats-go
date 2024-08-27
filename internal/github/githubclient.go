package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct {
	authToken string
}

// Keys are language names, values are number of bytes of code written
type RepoLanguages map[string]int

// Processed repository enriched with language data
type Repo struct {
	Name      string
	Languages map[string]int
	PushedAt  time.Time
}

func NewUnauthenticatedGithubClient() GithubClient {
	logger.Infof("Github client running in mode: unauthenticated\n")
	return GithubClient{}
}

func NewAuthenticatedGithubClient(authToken string) GithubClient {
	logger.Infof("Github client running in mode: authenticated\n")
	return GithubClient{authToken}
}

func (ghClient GithubClient) isAuthenticated() bool {
	return ghClient.authToken != ""
}

func (ghClient GithubClient) makeRequest(urlPath string) ([]byte, error) {
	req, err := http.NewRequest("GET", GithubAPIBaseURL+urlPath, nil)
	if err != nil {
		return nil, err
	}

	if ghClient.isAuthenticated() {
		req.Header.Set("Authorization", "Bearer "+ghClient.authToken)
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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
