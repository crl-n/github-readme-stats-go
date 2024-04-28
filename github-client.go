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

const repoCacheFilename = "cached_repos.json"

func CacheRepos(repos []Repo) {
	file, err := os.Create(repoCacheFilename)
	if err != nil {
		fmt.Println("Unable to create " + repoCacheFilename)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(repos)
	if err != nil {
		fmt.Println(err)
	}
}

func RetrieveCachedRepos() []Repo {
	file, err := os.Open(repoCacheFilename)
	if err != nil {
		fmt.Println("Unable to open " + repoCacheFilename)
		return nil
	}

	var repos []Repo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&repos)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return repos
}

func findRepo(repos []Repo, target RepoAPIResponse) (*Repo, bool) {
	for _, repo := range repos {
		if repo.Name == target.Name {
			fmt.Printf("Cache hit for %v\n", target.Name)
			return &repo, true
		}
	}
	fmt.Printf("Cache miss for %v\n", target.Name)
	return nil, false
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
			fmt.Printf("Using cached repo data for %v\n", rawRepo.Name)
			repos = append(repos, *cachedRepo)
		} else {
			fmt.Printf("Using new repo data for %v\n", rawRepo.Name)
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

func (ghClient GithubClient) GetLanguageStats() {
	repos, err := ghClient.GetUserRepos()
	if err != nil {
		return
	}

	fmt.Println(repos)
}
