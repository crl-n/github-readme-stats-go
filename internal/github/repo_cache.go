package github

import (
	"encoding/json"
	"os"

	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

const repoCacheFilename = "cached_repos.json"

func CacheRepos(repos []Repo) {
	file, err := os.Create(repoCacheFilename)
	if err != nil {
		logger.Errorf("Error creating file '%s': %s\n", repoCacheFilename, err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(repos)
	if err != nil {
		logger.Errorf("%v\n", err)
	}
}

func RetrieveCachedRepos() []Repo {
	file, err := os.Open(repoCacheFilename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Infof("Cached repos file '%s' not found. No cached repo data will be used.\n", repoCacheFilename)
		} else {
			logger.Errorf("Error opening file '%s': %s\n", repoCacheFilename, err)
		}
		return nil
	}

	var repos []Repo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&repos)
	if err != nil {
		logger.Errorf("%v\n", err)
		return nil
	}

	logger.Infof("Cached repos file '%s' found. Cached repo data will be used where possible.\n", repoCacheFilename)

	return repos
}

func findRepo(repos []Repo, target RawPublicRepo) (*Repo, bool) {
	for _, repo := range repos {
		if repo.Name == target.Name {
			return &repo, true
		}
	}
	return nil, false
}
