package github

import (
	"fmt"
	"time"

	. "github.com/crl-n/github-readme-stats-go/pkg/cache"
	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

type GithubService struct {
	ghClient  GithubClient
	repoCache *Cache[string, Repo]
}

const pubRepoCacheFilename = "cached_public_repos.json"

func NewGithubService(authToken string) *GithubService {
	if authToken == "" {
		return &GithubService{
			NewUnauthenticatedGithubClient(),
			NewCache[string, Repo](pubRepoCacheFilename),
		}
	}
	return &GithubService{
		NewAuthenticatedGithubClient(authToken),
		NewCache[string, Repo](pubRepoCacheFilename),
	}
}

// Fetches list of public repositories and enriches with language data for each
// repository.
func (repoService GithubService) GetPublicReposWithLanguages(githubHandle string) ([]Repo, error) {
	rawRepos, err := repoService.ghClient.GetPublicReposList(githubHandle)
	if err != nil {
		return nil, err
	}

	var repos []Repo

	for _, rawRepo := range rawRepos {
		rawRepoPushedAtTime, err := time.Parse(time.RFC3339, rawRepo.PushedAt)
		if err != nil {
			return nil, err
		}

		cacheKey := fmt.Sprintf("%s/%s", githubHandle, rawRepo.Name)
		cachedRepo, found := repoService.repoCache.Get(cacheKey)

		if found && cachedRepo.PushedAt.Equal(rawRepoPushedAtTime) {
			logger.Debugf("Cache hit for '%v', using cached repo data\n", rawRepo.Name)
			repos = append(repos, cachedRepo)
		} else {
			logger.Debugf("Cache miss for '%v', fetching language data\n", rawRepo.Name)
			repo, err := rawRepo.ToRepo(repoService.ghClient, githubHandle)
			if err != nil {
				return nil, err
			}
			repos = append(repos, repo)
		}
	}

	repoService.CachePublicReposData(githubHandle, repos)

	return repos, nil
}

func (repoService GithubService) CachePublicReposData(githubHandle string, repos []Repo) {
	var kvPairs []struct {
		Key   string
		Value Repo
	}

	for _, repo := range repos {
		kvPairs = append(kvPairs, struct {
			Key   string
			Value Repo
		}{
			Key:   fmt.Sprintf("%s/%s", githubHandle, repo.Name),
			Value: repo,
		})
	}

	repoService.repoCache.BulkSet(kvPairs)
}
