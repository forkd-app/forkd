package graph

import (
	"forkd/db"
	"forkd/services/auth"
	"forkd/services/email"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Queries *db.Queries
	Conn    *pgxpool.Pool
	Auth    auth.AuthService
	Email   email.EmailService
}
