package main

import (
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/githubclient"
	"github.com/crl-n/github-readme-stats-go/logger"
	"github.com/crl-n/github-readme-stats-go/stats"
	"github.com/crl-n/github-readme-stats-go/svg"
)

const (
	githubHandleEnvVar = "GITHUB_HANDLE"
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

	logger.Infof("Using Github handle '%s'\n", usedValue)
	return usedValue
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	githubHandle := getHandle()
	client := githubclient.New(githubHandle)

	switch os.Args[1] {
	case "lang":
		gen := svg.SVGGenerator{}

		repos, err := client.GetUserRepos()
		if err != nil {
			logger.Errorf("Error encountered while retrieving repositories: %v\n", err)
			return
		}

		languageStats := stats.NewLanguageStats(repos)

		for _, stat := range languageStats {
			logger.Infof("%v %v %v\n", stat.Language, stat.BytesOfCode, stat.Percentage)
		}

		gen.GenerateLangStatsCard(languageStats)
	case "gh":
		repos, err := client.GetUserRepos()
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
