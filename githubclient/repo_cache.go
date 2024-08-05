package githubclient

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/logger"
)

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

func findRepo(repos []Repo, target RawPublicRepo) (*Repo, bool) {
	for _, repo := range repos {
		if repo.Name == target.Name {
			logger.Debugf("Cache hit for %v\n", target.Name)
			return &repo, true
		}
	}
	logger.Debugf("Cache miss for %v\n", target.Name)
	return nil, false
}
