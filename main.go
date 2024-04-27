package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " svg\t\tGenerate demo svg file")
	fmt.Println(os.Args[0] + " gh\t\tGet and display Github stats")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "svg":
		GenerateTestSVG()
	case "gh":
		ghClient := GithubClient{"crl-n"}
		ghClient.GetUserData()
	default:
		usage()
	}
}
