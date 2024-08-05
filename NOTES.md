# Notes

- I've decided to focus solely on a language stats card for now, other cards may be added later.
- Currently, I temporarily cache the Github API data of individual repos to a file locally.
  - I do this to avoid issues with the Github API rate limits (60 requests per hour for unauthenticated users).
  - This method of caching needs to be replaced with something else or removed. A relational database could work, although I need to reflect on what will work best in a serverless environment.
- The language stats card needs some kind of gutters or padding.
- GithubClient.GetPublicReposWithLanguages does more than what might be expected from a method of GithubClient. This logic might be better to place somewhere else, for instance in a service. That way the concern of GithubClient would be more clear: fetching data from Github API.
