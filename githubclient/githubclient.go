package githubclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/crl-n/github-readme-stats-go/logger"
)

const GithubAPIBaseURL = "https://api.github.com"

type GithubClient struct {
	username string
}

func New(username string) GithubClient {
	return GithubClient{username}
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

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			fmt.Sprintf(
				"Request by GithubClient to '%s' failed with status %s",
				url,
				resp.Status,
			),
		)
	}

	return body, nil
}

// Raw public repository as serialized from JSON response from Github API
type RawPublicRepo struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
	PushedAt     string `json:"pushed_at"`
}

// Keys are language names, values are number of bytes of code written
type RepoLanguages map[string]int

// Processed repository enriched with language data
type Repo struct {
	Name      string
	Languages map[string]int
	PushedAt  time.Time
}

func (rawRepo RawPublicRepo) ToRepo(ghClient GithubClient) (Repo, error) {
	repoLanguages, err := ghClient.GetRepoLanguages(rawRepo.Name)
	if err != nil {
		return Repo{}, err
	}

	pushedAtTime, err := time.Parse(time.RFC3339, rawRepo.PushedAt)
	if err != nil {
		return Repo{}, err
	}

	repo := Repo{rawRepo.Name, repoLanguages, pushedAtTime}

	return repo, nil
}

// Fetches languages used in a repository. See:
// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repository-languages
func (ghClient GithubClient) GetRepoLanguages(repo string) (RepoLanguages, error) {
	body, err := ghClient.makeRequest(
		GithubAPIBaseURL + "/repos/" + ghClient.username + "/" + repo + "/languages",
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
func (ghClient GithubClient) GetPublicReposList() ([]RawPublicRepo, error) {
	body, err := ghClient.makeRequest(GithubAPIBaseURL + "/users/" + ghClient.username + "/repos")
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

// Fetches list of public repositories and enriches with language data for each
// repository.
func (ghClient GithubClient) GetPublicReposWithLanguages() ([]Repo, error) {
	rawRepos, err := ghClient.GetPublicReposList()
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
