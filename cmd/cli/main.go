package main

import (
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/internal/cards"
	"github.com/crl-n/github-readme-stats-go/internal/github"
	"github.com/crl-n/github-readme-stats-go/internal/stats"
	"github.com/crl-n/github-readme-stats-go/pkg/env"
	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " gh [handle]\t\tGet and display Github stats")
	fmt.Println(os.Args[0] + " lang [handle]\t\tGenerate most used languages card")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	authToken := env.GetAuthToken()
	githubHandle := env.GetHandle()
	service := github.NewGithubService(authToken)

	switch os.Args[1] {
	case "lang":
		repos, err := service.GetPublicReposWithLanguages(githubHandle)
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
	case "gh":
		repos, err := service.GetPublicReposWithLanguages(githubHandle)
		if err != nil {
			logger.Errorf("Error encountered while retrieving repositories: %v\n", err)
			return
		}

		for _, repo := range repos {
			logger.Infof("%v\n", repo)
		}
	default:
		usage()
	}
}
