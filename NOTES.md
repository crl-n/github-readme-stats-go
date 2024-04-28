# Notes

- I've decided to focus solely on a language stats card for now, other cards may be added later.
- Currently, I temporarily cache the Github API data of individual repos to a file locally. I do this to avoid issues with the Github API rate limits (60 requests per hour for unauthenticated users).
- GithubClient now contains some aggregation logic that may be better placed elsewhere, perhaps somekind of aggregator struct like LangStatsAggregator.
- 
