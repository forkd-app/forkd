package model

import (
	"forkd/db"
)

func RecipeFromDBType(result db.Recipe) *Recipe {
	// Map to model.Recipe type
	recipe := Recipe{
		ID:                 int(result.ID),
		Slug:               result.Slug,
		InitialPublishDate: result.InitialPublishDate.Time,
		Private:            result.Private,
	}

	return &recipe
}

func UserFromDBType(result db.User) *User {
	// Map to model.User type
	user := User{
		ID:          int(result.ID),
		Email:       result.Email,
		JoinDate:    result.JoinDate.Time,
		DisplayName: result.DisplayName,
		UpdatedAt:   result.UpdatedAt.Time,
	}

	return &user
}
