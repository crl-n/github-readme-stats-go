package github

import "time"

// Raw public repository as serialized from JSON response from Github API
type RawPublicRepo struct {
	Name         string `json:"name"`
	LanguagesUrl string `json:"languages_url"`
	PushedAt     string `json:"pushed_at"`
}

// Enriches raw public repo with language data
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
