package github

import (
	"time"

	"github.com/crl-n/github-readme-stats-go/logger"
)

type GithubService struct {
	ghClient GithubClient
}

func NewGithubService() GithubService {
	return GithubService{NewGithubClient()}
}

// Fetches list of public repositories and enriches with language data for each
// repository.
func (repoService GithubService) GetPublicReposWithLanguages(githubHandle string) ([]Repo, error) {
	rawRepos, err := repoService.ghClient.GetPublicReposList(githubHandle)
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
			logger.Debugf("Cache hit for '%v', using cached repo data\n", rawRepo.Name)
			repos = append(repos, *cachedRepo)
		} else {
			logger.Debugf("Cache miss for '%v', fetching language data\n", rawRepo.Name)
			repo, err := rawRepo.ToRepo(repoService.ghClient, githubHandle)
			if err != nil {
				return nil, err
			}
			repos = append(repos, repo)
		}
	}

	CacheRepos(repos)

	return repos, nil
}
