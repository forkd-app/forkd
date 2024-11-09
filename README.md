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

- Docker
- Probably Node >= v20
- PNPM
- Go >= v1.22

### Starting the application services

To spin up the application services, you just need to make sure you have the dependencies installed and run `docker-compose up -d` from the root directory. This will start all the services specifies in the `docker-compose.yml`.

### Scripts

We have a couple scripts setup inside the `docker-compose.yml`. Here is an overview of them:

#### GraphQL

We use [gqlgen](https://gqlgen.com/) for generating Go types from our GraphQL [schema](api/graph/schema). You can access this feature using the `gqlgen` compose service in the `scripts` profile. Here are some examples:

- `docker-compose --profile "scripts" run --rm gqlgen "generate"`
  - This generates the GraphQL types for our API from the schema files.

#### DB Migrations

We use [geni](https://github.com/emilpriver/geni) for our DB migration tool. This manages the migrations in the [migrations](db/migrations) folder. Some examples of using this from docker-compose:

- `docker-compose --profile "scripts" run --rm migrate new`
  - This creates a new migration file
- `docker-compose --profile "scripts" run --rm migrate up`
  - This applies any needed migrations
- `docker-compose --profile "scripts" run --rm migrate down`
  - Roll back applied migrations

#### DB Queries

We use [sqlc](https://github.com/sqlc-dev/sqlc) to generate typesafe `go` code based on our `sql` queries. This uses the queries in the [queries](db/queries/) folder. Some examples of using this from docker-compose:

- `docker-compose --profile "scripts" run --rm sqlgen generate`
  - This generates go files for our queries
