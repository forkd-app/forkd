package graph

import (
	"forkd/services/auth"
	"forkd/services/email"
	"forkd/services/recipe"
	"forkd/services/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService   auth.AuthService
	EmailService  email.EmailService
	RecipeService recipe.RecipeService
	UserService   user.UserService
}
