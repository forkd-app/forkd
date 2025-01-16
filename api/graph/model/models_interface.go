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

func RevisionFromDBType(result db.RecipeRevision) *RecipeRevision {
	revision := RecipeRevision{
		ID:                int(result.ID),
		RecipeDescription: &result.RecipeDescription.String,
		ChangeComment:     &result.ChangeComment.String,
		Title:             result.Title,
		PublishDate:       result.PublishDate.Time,
	}

	return &revision
}

func ListIngredientsFromDBType(results []db.RecipeIngredient) []*RecipeIngredient {
	ingredients := make([]*RecipeIngredient, len(results))

	for i, result := range results {
		ingredients[i] = &RecipeIngredient{
			ID:       int(result.ID),
			Quantity: float64(result.Quantity),
			Comment:  &result.Comment.String,
		}
	}

	return ingredients
}

func ListStepsFromDBType(results []db.RecipeStep) []*RecipeStep {
	steps := make([]*RecipeStep, len(results))

	for i, result := range results {
		steps[i] = &RecipeStep{
			ID:      int(result.ID),
			Content: result.Content,
			Index:   int(result.Index),
		}
	}

	return steps

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
