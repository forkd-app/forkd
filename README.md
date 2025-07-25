# Forkd

This is a monorepo with the following packages:

- `API`
  - path
    - `./api`
  - [Go](https://go.dev/)
  - [GraphQL](https://graphql.org/)
  - [PostgreSQL](https://www.postgresql.org/)
  - [Minio](https://min.io/)
- `Web`
  - path
    - `./web`
  - [Remix](https://remix.run/)
  - [Mantine](https://mantine.dev/)

## Development

I've tried to make this pretty easy to develop with. We use Docker and Docker Compose to keep everything wrapped up nice and tidy.

### Dependencies

- [Docker](https://docs.docker.com/engine/install/)
  - For running all the services together and reducing the required deps to install
- [Task](https://taskfile.dev/installation/)
  - For running commands in a simpler fashion
- [Node](https://nodejs.org/en/download/package-manager/current) >= v20
  - Frontend
  - Prettier
  - GraphQL linting
- [PNPM](https://pnpm.io/installation)
  - Package manager if you want to run the app without docker
- [Go](https://go.dev/doc/install) >= v1.22
  - Backend

### Setup

To start things off, we should run `task bootstrap`. This does a couple things:

- Sets up the `.env` files in for the apps
- Pulls all required Docker images
- Builds the local docker images
- Verifies dependencies
- Sets up git hooks

### Starting the application services

To spin up the application services, you just need to make sure you have the dependencies installed and run `task up` from the root directory. This will start all the services specifies in the `docker-compose.yml`. You can run `task watch` to start the services in `watch` mode. This will rebuild/restart the services when files change.

To bring the services down run `task down`.

#### Service Ports

The services run on the following ports:

| Service       | Port |
| ------------- | ---- |
| Web           | 3000 |
| API           | 8000 |
| PgWeb         | 8081 |
| Minio Console | 9001 |
| Postgres      | 5432 |
| Minio         | 9000 |

#### Logging In/Sign Up

When logging in or signing up by default we don't send the magic link emails. To complete the auth flow do the following:

- Either log in or sign up until you get to the "Please check your email screen"
- Look for the line in your terminal that looks like this:
  - `MAGIC LINK CODE: ${CODE}`
  - _**note**: If running in detached mode, like running `task up`, run `docker compose logs api`_
- In your browser go to `http://localhost:3000/auth/validate?code=${CODE}`

🎉🎉🎉

### Scripts

We have a couple scripts setup inside the [docker-compose.yml](./docker-compose.yml). I have setup additional shortcuts from them in the [taskfile](./Taskfile.yml)

#### GraphQL

We use [gqlgen](https://gqlgen.com/) for generating Go types from our GraphQL [schema](api/graph/schema). You can access this feature using the `gqlgen` compose service in the `scripts` profile. Here are some examples:

- `task gqlgen`
  - This generates the GraphQL types for our API from the schema files.

#### DB Migrations

We use [go migrate](https://github.com/golang-migrate/migrate) for our DB migration tool. This manages the migrations in the [migrations](db/migrations) folder. Some examples of using this from docker-compose:

- `task migrate-new`
  - This creates a new migration file
- `task migrate-up`
  - This applies any needed migrations
- `task migrate-down`
  - Roll back applied migrations

#### DB Queries

We use [sqlc](https://github.com/sqlc-dev/sqlc) to generate typesafe `go` code based on our `sql` queries. This uses the queries in the [queries](db/queries/) folder. Some examples of using this from docker-compose:

- `task sqlc-gen`
  - This generates go files for our queries

#### DB Seed

- `task seed`
  - Generate and insert random data into the database
- task reset
  - Alias for `task drop`, `task migrate-up`, and `task seed`

#### Docker Commands

- task up
  - Bring up all the main services in detached mode
- task down
  - Bring down all the main services
- task watch
  - Bring up all the main services and watches for changes
- task build
  - Build all local docker images
- task pull
  - Pull all the required docker images

#### Linting

- task lint:go
  - Run golangci-lint
- task lint:go:fix
  - Run golangci-lint with --fix flag
- task lint:ts
  - Run ESLint
- task lint:ts:fix
  - Run ESLint with --fix flag
- task lint
  - Alias for `task lint:go` and `task lint:ts`
- task lint:fix
  - Alias for `task lint:go:fix` and `task lint:ts:fix`

#### Formatting

- task gofmt
  - Run gofmt
- task gofmt:check
  - Run gofmt and error if anything isn't formatted correctly
- task prettier
  - Run prettier
- task prettier:check
  - Run prettier and error if anything isn't formatted correctly
- task format
  - Alias for `task gofmt` and `task prettier`
- task format:check
  - Alias for `task gofmt:check` and `task prettier:check`
