# Notes

- I've decided to focus solely on a language stats card for now, other cards may be added later.
- Currently, I temporarily cache the Github API data of individual repos to a file locally.
  - I do this to avoid issues with the Github API rate limits (60 requests per hour for unauthenticated users).
  - This method of caching needs to be replaced with something else or removed. A relational database could work, although I need to reflect on what will work best in a serverless environment.
