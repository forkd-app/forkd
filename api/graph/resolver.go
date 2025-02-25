package graph

import (
	"forkd/db"
	"forkd/services/auth"
	"forkd/services/email"
	"forkd/services/recipe"
	"forkd/services/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Queries       *db.Queries
	Conn          *pgxpool.Pool
	AuthService   auth.AuthService
	EmailService  email.EmailService
	RecipeService recipe.RecipeService
	UserService   user.UserService
}
