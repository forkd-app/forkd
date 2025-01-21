package model

import (
	"forkd/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapStringToPgUuid(str string) (pgtype.UUID, error) {
	id, err := uuid.Parse(str)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}, nil
}

func MapPgUuidToString(id pgtype.UUID) (string, error) {
	goUuid, err := uuid.FromBytes(id.Bytes[:])
	if err != nil {
		return "", err
	}

	return goUuid.String(), nil
}

func RecipeFromDBType(result db.Recipe) *Recipe {
	id, err := MapPgUuidToString(result.ID)
	if err != nil {
		return nil
	}
	recipe := Recipe{
		ID:                 id,
		Slug:               result.Slug,
		InitialPublishDate: result.InitialPublishDate.Time,
		Private:            result.Private,
	}

	return &recipe
}

func UserFromDBType(result db.User) *User {
	id, err := MapPgUuidToString(result.ID)
	if err != nil {
		return nil
	}
	user := User{
		ID:          id,
		Email:       result.Email,
		JoinDate:    result.JoinDate.Time,
		DisplayName: result.DisplayName,
		UpdatedAt:   result.UpdatedAt.Time,
	}

	return &user
}

func RevisionFromDBType(result db.RecipeRevision) *RecipeRevision {
	id, err := MapPgUuidToString(result.ID)
	if err != nil {
		return nil
	}
	revision := RecipeRevision{
		ID:                id,
		RecipeDescription: IfValidString(result.RecipeDescription),
		ChangeComment:     IfValidString(result.ChangeComment),
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

func IfValidString(text pgtype.Text) *string {
	if text.Valid {
		return &text.String
	}
	return nil
}
