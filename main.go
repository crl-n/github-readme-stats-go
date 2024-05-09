package main

import (
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/githubclient"
	"github.com/crl-n/github-readme-stats-go/logger"
	"github.com/crl-n/github-readme-stats-go/stats"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " svg\t\tGenerate demo svg file")
	fmt.Println(os.Args[0] + " gh\t\tGet and display Github stats")
	fmt.Println(os.Args[0] + " lang\t\tGenerate most used languages card")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	client := githubclient.New("crl-n")

	switch os.Args[1] {
	case "lang":
		gen := SVGGenerator{}

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
	case "svg":
		GenerateTestSVG()
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
