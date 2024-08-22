# Github Readme Stats Go

Inspired by [Anurag Hazra's github-readme-stats](https://github.com/anuraghazra/github-readme-stats) and similar projects.

I was curious how github-readme-stats was able to show dynamic content in the static context of a Github README. If you're curious about how this works I recommend reading [Bohdan Liashenko's blog post about it](https://codecrumbs.io/library/github-readme-stats). This project is my attempt at achieving the same (or something similar) using Go with **zero dependencies**. Most of all this project is my outlet for learning Go on the side of my real work.

At this point the project is not set up to work exactly like github-readme-stats does, as it is early in development. There is no server yet as I opted to develop the core functionality as a CLI. The plan is to turn it into a server that can be self-hosted once the core functionality is implemented.

I hope you enjoy the project!

## Use
Build the project using `go build`. Use compiled binary to generate SVG card with stats of a Github user:
```
./github-readme-stats-go lang crl-n
```
This will generate a langs.svg file that contains the top languages used in all public repositories of the user.

The github user can also be specified with an environment variable, in which case username argument can be omitted:
```
export GITHUB_USERNAME=crl-n
./github-readme-stats-go lang
```

## GitHub API Authentication
You may choose to authenticate to the GitHub API. This allows for more requests per hour. To do this, set `GITHUB_AUTH_TOKEN` environment variable.
```
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
