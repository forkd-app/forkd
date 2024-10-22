# Go Version
FROM golang:1.23-alpine

# Basic setup of the container
RUN mkdir /app
COPY .. /app
WORKDIR /app
RUN go mod download
RUN go mod vendor

# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built
ENTRYPOINT ["go", "run", "-mod=vendor", "github.com/99designs/gqlgen"]