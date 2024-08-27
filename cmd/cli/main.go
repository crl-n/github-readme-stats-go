package main

import (
	"fmt"
	"os"

	"github.com/crl-n/github-readme-stats-go/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	cli := cli.NewCLI()

	switch os.Args[1] {
	case "lang":
		cli.HandleLang()
	case "gh":
		cli.HandleGH()
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0] + " gh [handle]\t\tGet and display Github stats")
	fmt.Println(os.Args[0] + " lang [handle]\t\tGenerate most used languages card")
}
