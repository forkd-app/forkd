# Forkd:

This is a monorepo with the following packages:

- `API`
  - path
    - `./api`
  - Go app using GraphQL and PostgreSQL
- `Web`
  - path
    - `./web`
  - Remix app

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

### Scripts

We have a couple scripts setup inside the [docker-compose.yml](./docker-compose.yml). I have setup additional shortcuts from them in the [taskfile](./Taskfile.yml):

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
