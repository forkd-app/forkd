package main

import (
	"forkd/db"
	"forkd/graph"
	"forkd/services/auth"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// TODO: Make this an env var
	queries, _, err := db.GetQueriesWithConnection("postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable")
	if err != nil || queries == nil {
		panic("Unable to connect to db")
	}

	authService := auth.New(queries)

	srvConf := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Queries: *queries, Auth: authService}, Directives: graph.DirectiveRoot{Auth: graph.AuthDirective(authService)}})
	srv := handler.NewDefaultServer(srvConf)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authService.SessionWrapper(srv.ServeHTTP))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
