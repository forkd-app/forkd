# Forkd

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
