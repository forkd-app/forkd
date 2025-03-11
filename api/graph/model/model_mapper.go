package model

import (
	"forkd/db"

	"github.com/jackc/pgx/v5/pgtype"
)

func RecipeFromDBType(result db.Recipe) *Recipe {
	recipe := Recipe{
		ID:                 result.ID.Bytes,
		Slug:               result.Slug,
		InitialPublishDate: result.InitialPublishDate.Time,
		Private:            result.Private,
		ForkedFrom: &RecipeRevision{
			ID: result.ForkedFrom.Bytes,
		},
		FeaturedRevision: &RecipeRevision{
			ID: result.FeaturedRevision.Bytes,
		},
		Author: &User{
			ID: result.AuthorID.Bytes,
		},
	}

	return &recipe
}

func UserFromDBType(result db.User) *User {
	user := User{
		ID:          result.ID.Bytes,
		Email:       result.Email,
		JoinDate:    result.JoinDate.Time,
		DisplayName: result.DisplayName,
		UpdatedAt:   result.UpdatedAt.Time,
		Photo:       IfValidString(result.Photo),
	}

	return &user
}

func RevisionFromDBType(result db.RecipeRevision) *RecipeRevision {
	revision := RecipeRevision{
		ID:                result.ID.Bytes,
		RecipeDescription: IfValidString(result.RecipeDescription),
		ChangeComment:     IfValidString(result.ChangeComment),
		Title:             result.Title,
		PublishDate:       result.PublishDate.Time,
		Recipe: &Recipe{
			ID: result.RecipeID.Bytes,
		},
		Parent: &RecipeRevision{
			ID: result.ParentID.Bytes,
		},
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
			Revision: &RecipeRevision{
				ID: result.RevisionID.Bytes,
			},
			Ingredient: &Ingredient{
				ID: int(result.IngredientID),
			},
			Unit: &MeasurementUnit{
				ID: int(result.MeasurementUnitID),
			},
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
			Revision: &RecipeRevision{
				ID: result.RevisionID.Bytes,
			},
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

func IfValidString(text pgtype.Text) *string {
	if text.Valid {
		return &text.String
	}
	return nil
}
