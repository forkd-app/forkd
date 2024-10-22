# Forkd API

This is our main backend API for Forkd. We use the following technologies:

- `Go`
  - This is the programming language the API is written in.
- `GraphQL`
  - This is the main API surface.
  - We use `gqlgen` for auto generating resolver stubs from our GraphQL schema.
- `PostgreSQL`
  - Our main database.
- `Docker`
  - We use this to try and keep our dev environments with as little setup as possible.
- `Docker Compose`
  - Just an easy way to orchestrate a couple of docker containers for development.
  - This lets us easily start the API, the DB, and any other services we need with a single command.
