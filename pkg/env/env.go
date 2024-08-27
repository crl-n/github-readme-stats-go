package env

import (
	"os"

	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

const (
	githubHandleEnvVar    = "GITHUB_HANDLE"
	githubAuthTokenEnvVar = "GITHUB_AUTH_TOKEN"
)

// Gets Github handle to be used. If no Github handle provided as command line
// argument, environment variable is used. If no handle provided, exit gracefully.
func GetHandle() string {
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

func GetAuthToken() string {
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
