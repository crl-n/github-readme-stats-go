package cli

import (
	"github.com/crl-n/github-readme-stats-go/internal/cards"
	"github.com/crl-n/github-readme-stats-go/internal/github"
	"github.com/crl-n/github-readme-stats-go/internal/stats"
	"github.com/crl-n/github-readme-stats-go/pkg/env"
	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

type CLI struct {
	AuthToken    string
	GithubHandle string
	Service      *github.GithubService
}

func NewCLI() *CLI {
	authToken := env.GetAuthToken()
	githubHandle := env.GetHandle()
	service := github.NewGithubService(authToken)

	return &CLI{
		AuthToken:    authToken,
		GithubHandle: githubHandle,
		Service:      service,
	}
}

func (app *CLI) HandleLang() {
	repos, err := app.Service.GetPublicReposWithLanguages(app.GithubHandle)
	if err != nil {
		logger.Errorf("Error encountered while retrieving repositories: %v\n", err)
		return
	}

	languageStats := stats.NewLanguageStats(repos)

	for _, stat := range languageStats {
		logger.Infof("%v %v %v\n", stat.Language, stat.BytesOfCode, stat.Percentage)
	}

	langStatsCard := cards.NewLanguageStatsCard(languageStats)
	langStatsCard.GenerateSVGFile()
}

func (app *CLI) HandleGH() {
	repos, err := app.Service.GetPublicReposWithLanguages(app.GithubHandle)
	if err != nil {
		logger.Errorf("Error encountered while retrieving repositories: %v\n", err)
		return
	}

	for _, repo := range repos {
		logger.Infof("%v\n", repo)
	}
}
