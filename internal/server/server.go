package server

import (
	"fmt"
	"net/http"

	"github.com/crl-n/github-readme-stats-go/internal/github"
	"github.com/crl-n/github-readme-stats-go/pkg/env"
	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

const (
	port = "8080"
)

type Server struct {
	AuthToken     string
	GithubService *github.GithubService
}

func NewServer() *Server {
	authToken := env.GetAuthToken()

	return &Server{
		AuthToken: authToken,
	}
}

func (server *Server) Start() {
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/langs/", handleLangs)

	logger.Infof("Starting server on :%v, listening...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		logger.Errorf("Error starting server: %v\n", err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func handleLangs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Placeholder"))
}
