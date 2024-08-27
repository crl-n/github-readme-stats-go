package main

import "github.com/crl-n/github-readme-stats-go/internal/server"

func main() {
	server := server.NewServer()

	server.Start()
}
