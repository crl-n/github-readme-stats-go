package main

import (
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/cards"
	"github.com/crl-n/github-readme-stats-go/github"
	"github.com/crl-n/github-readme-stats-go/logger"
	"github.com/crl-n/github-readme-stats-go/stats"
)

const (
	githubHandleEnvVar    = "GITHUB_HANDLE"
	githubAuthTokenEnvVar = "GITHUB_AUTH_TOKEN"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " gh [handle]\t\tGet and display Github stats")
	fmt.Println(os.Args[0] + " lang [handle]\t\tGenerate most used languages card")
}

// Gets Github handle to be used. If no Github handle provided as command line
// argument, environment variable is used. If no handle provided, exit gracefully.
func getHandle() string {
	var usedValue string

	if len(os.Args) == 3 {
		argValue := os.Args[2]
		usedValue = argValue
	} else {
		envValue := os.Getenv(githubHandleEnvVar)

		if envValue == "" {
			logger.Errorf(
				"Error: No Github handle provided. Provide handle as argument or "+
					"set environment variable '%s'\n", githubHandleEnvVar,
			)
			os.Exit(1)
		}
		usedValue = envValue
	}

	logger.Infof("Using target Github handle '%s'\n", usedValue)
	return usedValue
}

func getAuthToken() string {
	envValue := os.Getenv(githubAuthTokenEnvVar)

	if envValue == "" {
		logger.Infof(
			"No Github auth token provided. Requests to GitHub API will "+
				"be unauthenticated, which limits requests per hour. To "+
				"authenticate set environment variable '%s'\n", githubAuthTokenEnvVar,
		)
	}

	return envValue
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	authToken := getAuthToken()
	githubHandle := getHandle()
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
