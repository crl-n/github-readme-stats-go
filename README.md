# Github Readme Stats Go

Inspired by [Anurag Hazra's github-readme-stats](https://github.com/anuraghazra/github-readme-stats) and similar projects.

I was curious how github-readme-stats was able to show dynamic content in the static context of a Github README. If you're curious about how this works I recommend reading [Bohdan Liashenko's blog post about it](https://codecrumbs.io/library/github-readme-stats). This project is my attempt at achieving the same using Go. I hope you enjoy the project!

## Logging
Log level is configurable by the LOG_LEVEL environment variable. Available log levels:
```
LOG_LEVEL=DEBUG
LOG_LEVEL=INFO
LOG_LEVEL=ERROR
```
Default log level is `INFO`.