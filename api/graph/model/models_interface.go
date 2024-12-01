package model

import "forkd/db"

func RecipeFromDBType(result db.Recipe) *Recipe {
	// Handle nullable "forkedFrom" value
	var forkedFrom *int = nil
	if result.ForkedFrom.Valid {
		val := int(result.ForkedFrom.Int64)
		forkedFrom = &val
	}
	// Map to model.Recipe type
	recipe := Recipe{
		ID:                 int(result.ID),
		Slug:               result.Slug,
		ForkedFrom:         forkedFrom,
		InitialPublishDate: result.InitialPublishDate.Time.String(),
		Description:        result.Description.String,
	}

	return &recipe
}

func UserFromDBType(result db.User) *User {
	// Map to model.User type
	user := User{
		ID:       int(result.ID),
		Email:    result.Email,
		JoinDate: result.JoinDate.Time.String(),
		Username: result.Username,
	}

	return &user
}