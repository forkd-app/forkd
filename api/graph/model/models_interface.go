package model

import (
	"forkd/db"

	"github.com/jackc/pgx/v5/pgtype"
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

func RevisionFromDBType(result db.RecipeRevision) *RecipeRevision {
	revision := RecipeRevision{
		ID:                int(result.ID),
		RecipeDescription: IfValidString(result.RecipeDescription),
		ChangeComment:     IfValidString(result.ChangeComment),
		Title:             result.Title,
		PublishDate:       result.PublishDate.Time,
	}
	return &revision
}

func MeasurementUnitFromDBType(result db.MeasurementUnit) *MeasurementUnit {
	unit := MeasurementUnit{
		ID:          int(result.ID),
		Name:        result.Name,
		Description: &result.Description.String,
	}
	return &unit
}

func IngredientFromDBType(result db.Ingredient) *Ingredient {
	ingredient := Ingredient{
		ID:          int(result.ID),
		Name:        result.Name,
		Description: &result.Description.String,
	}
	return &ingredient
}

func IfValidString(text pgtype.Text) *string {
	if text.Valid {
		return &text.String
	}
	return nil
}
