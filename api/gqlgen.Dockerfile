# Go Version
FROM golang:1.23-alpine

# Basic setup of the container
RUN mkdir /app
COPY .. /app
WORKDIR /app
RUN go mod tidy
ENTRYPOINT ["go", "run", "github.com/99designs/gqlgen"]
