# Github Readme Stats Go

Inspired by [Anurag Hazra's github-readme-stats](https://github.com/anuraghazra/github-readme-stats) and similar projects.

I was curious how github-readme-stats was able to show dynamic content in the static context of a Github README. If you're curious about how this works I recommend reading [Bohdan Liashenko's blog post about it](https://codecrumbs.io/library/github-readme-stats). This project is my attempt at achieving the same (or something similar) using Go with **zero dependencies**. Most of all this project is my outlet for learning Go on the side of my real work.

There are two build targets that may be compiled: a HTTP-server and a CLI. The server is the main build target intended for production use. The CLI is intended for development use only. The CLI makes it easy to test the core functionality (SVG-generation etc.) without the need of making HTTP requests.

At this point the project is not set up to work exactly like github-readme-stats does. I've opted to target VM/container compute instead of the serverless approach used by Hazra. I also chose to implement only the top languages card for now. My reasoning is that I want to achieve one fully functioning card before considering implementing additional cards.

I hope you enjoy the project!

## Server

### Build
To build the server use command:
```sh
go build -o bin/server cmd/server/main.go
```

### Usage
To start the server use command:
```sh
./bin/server
```
The server should start listening for incoming requests on port 8080.

## CLI

### Build
To build the CLI use command:
```sh
go build -o bin/cli cmd/cli/main.go
```

### Usage
After building, use compiled binary to generate SVG card with stats of a Github user:
```sh
./bin/cli lang crl-n
```
This will generate a langs.svg file that contains the top languages used in all public repositories of the user.

The github user can also be specified with an environment variable, in which case username argument can be omitted:
```sh
export GITHUB_USERNAME=crl-n
./bin/cli lang
```

### GitHub API Authentication
You may choose to authenticate to the GitHub API. This allows for more requests per hour. To do this, set `GITHUB_AUTH_TOKEN` environment variable.
```sh
export GITHUB_AUTH_TOKEN=my-token
```

## Logging
Log level is configurable by the LOG_LEVEL environment variable. Available log levels:
```
LOG_LEVEL=DEBUG
LOG_LEVEL=INFO
LOG_LEVEL=ERROR
```
Default log level is `INFO`.
