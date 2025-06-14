package main

import (
	"encoding/json"
	"fmt"
	"forkd/db"
	"forkd/graph"
	"forkd/services/auth"
	"forkd/services/email"
	"forkd/services/object_storage"
	"forkd/services/recipe"
	"forkd/services/user"
	"forkd/util"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
)

type uploadType int

const (
	profilePhoto uploadType = iota
	revisionPhoto
	stepPhoto
)

type GetUploadUrlBody struct {
	UploadType uploadType
	StepIndex  int       `json:"stepIndex,omitempty"`
	RandId     uuid.UUID `json:"randId,omitempty"`
}

func main() {
	util.InitEnv()
	env := util.GetEnv()
	dbConnStr := env.GetDbConnStr()
	port := env.GetPort()

	queries, conn, err := db.GetQueriesWithConnection(dbConnStr)
	if err != nil || queries == nil {
		panic(fmt.Errorf("unable to connect to db: %w", err))
	}

	emailService := email.New()
	authService := auth.New(queries, conn, emailService)
	photoService := object_storage.New("forkd")
	recipeService := recipe.New(queries, conn, authService, photoService)
	userService := user.New(queries, authService)

	// TODO: We should do a refactor here, it's getting pretty cluttered (Mostly my fault lol)
	srvConf := graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			AuthService:   authService,
			EmailService:  emailService,
			RecipeService: recipeService,
			UserService:   userService,
		},
		Directives: graph.DirectiveRoot{
			Auth: graph.AuthDirective(authService),
		},
	})
	srv := handler.NewDefaultServer(srvConf)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authService.SessionWrapper(srv.ServeHTTP))
	// TODO: Move this into a package, kinda ugly here lol
	http.Handle("POST /get-upload-url", authService.SessionWrapper(func(w http.ResponseWriter, r *http.Request) {
		ctx := authService.GetUserSessionAndSetOnCtx(r.Context())
		user, _ := authService.GetUserSessionFromCtx(ctx)
		if user == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var body GetUploadUrlBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "invalid body", 400)
			return
		}
		var filename string
		switch body.UploadType {
		case profilePhoto:
			filename = fmt.Sprintf("%s/profilephoto", uuid.UUID(user.ID.Bytes).String())
		case revisionPhoto:
			if body.RandId == uuid.Nil {
				http.Error(w, "invalid rand id", 400)
				return
			}
			filename = fmt.Sprintf("%s/%s/revisionphoto", uuid.UUID(user.ID.Bytes).String(), body.RandId.String())
		case stepPhoto:
			if body.RandId == uuid.Nil {
				http.Error(w, "invalid rand id", 400)
				return
			}
			if body.StepIndex == 0 {
				http.Error(w, "invalid step index", 400)
				return
			}
			filename = fmt.Sprintf("%s/%s/step-%d", uuid.UUID(user.ID.Bytes).String(), body.RandId.String(), body.StepIndex)
		default:
			http.Error(w, "invalid upload type", 400)
		}
		url, err := photoService.GetUploadUrl(ctx, filename, time.Minute*3)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		_, err = w.Write([]byte(url))
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
