package main

import (
	"fmt"
	"forkd/db"
	"forkd/graph"
	"forkd/services/auth"
	"forkd/services/email"
	"forkd/util"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	util.InitEnv()
	env := util.GetEnv()
	dbConnStr := env.GetDbConnStr()
	port := env.GetPort()

	queries, conn, err := db.GetQueriesWithConnection(dbConnStr)
	if err != nil || queries == nil {
		panic(fmt.Errorf("Unable to connect to db: %w", err))
	}

	emailService := email.New()
	authService := auth.New(queries, conn)

	// TODO: We should do a refactor here, it's getting pretty cluttered (Mostly my fault lol)
	srvConf := graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Queries: queries,
			Conn:    conn,
			Auth:    authService,
			Email:   emailService,
		},
		Directives: graph.DirectiveRoot{
			Auth: graph.AuthDirective(authService),
		},
	})
	srv := handler.NewDefaultServer(srvConf)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authService.SessionWrapper(srv.ServeHTTP))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
