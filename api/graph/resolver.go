package graph

import (
	"forkd/db"
	"forkd/services/auth"
	"forkd/services/email"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Queries db.Queries
	Auth    auth.AuthService
	Email   email.EmailService
}
