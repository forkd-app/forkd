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

func UserFromDBType[T any](result T) *User {
	switch v := any(result).(type) {
	case db.User:
		return &User{
			ID:          int(v.ID),
			Email:       v.Email,
			JoinDate:    v.JoinDate.Time,
			DisplayName: v.DisplayName,
			UpdatedAt:   v.UpdatedAt.Time,
		}
	case db.GetRecipeWithAuthorByIdRow:
		return &User{
			ID:          int(v.UserID),
			Email:       v.UserEmail,
			JoinDate:    v.UserJoinDate.Time,
			DisplayName: v.UserDisplayName,
			UpdatedAt:   v.UserUpdatedAt.Time,
		}
	default:
		return nil
	}
}

func RevisionFromDBType[T any](result T) *RecipeRevision {
	switch v := any(result).(type) {
	case db.RecipeRevision:
		return &RecipeRevision{
			ID:                int(v.ID),
			RecipeDescription: ifValidString(v.RecipeDescription),
			ChangeComment:     ifValidString(v.ChangeComment),
			Title:             v.Title,
			PublishDate:       v.PublishDate.Time,
		}
	case db.GetRecipeWithForkedFromByIdRow:
		if !v.ForkedRevisionID.Valid {
			return nil
		}
		return &RecipeRevision{
			ID:                int(v.ForkedRevisionID.Int64),
			RecipeDescription: ifValidString(v.ForkedRecipeDescription),
			ChangeComment:     ifValidString(v.ForkedChangeComment),
			Title:             v.ForkedTitle.String,
			PublishDate:       v.ForkedPublishDate.Time,
		}
	default:
		return nil
	}
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

func ifValidString(text pgtype.Text) *string {
	if text.Valid {
		return &text.String
	}
	return nil
}
