# Notes

- I've decided to focus solely on a language stats card for now, other cards may be added later.
- Currently, I temporarily cache the Github API data of individual repos to a file locally.
  - I do this to avoid issues with the Github API rate limits (60 requests per hour for unauthenticated users).
  - This method of caching needs to be replaced with something else or removed. A relational database could work, although I need to reflect on what will work best in a serverless environment.
- Currently the github client and language stats are tightly coupled, which makes it hard to turn them into packages without causing circular dependencies and other issues. It would be beneficial to decouple the two and possible turn them into their own packages.
