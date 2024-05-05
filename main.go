package main

import (
	"fmt"
	"os"
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

	ghClient := GithubClient{"crl-n"}

	switch os.Args[1] {
	case "lang":
		gen := SVGGenerator{}
		stats := ghClient.GetLanguageStats()
		gen.GenerateLangStatsCard(stats)
	case "svg":
		GenerateTestSVG()
	case "gh":
		ghClient.GetLanguageStats()
	default:
		usage()
	}
}
