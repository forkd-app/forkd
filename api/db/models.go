// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Ingredient struct {
	Name        string
	Description pgtype.Text
}

type IngredientTag struct {
	Ingredient string
	Tag        string
}

type LinkedRecipe struct {
	FromRecipeID int64
	ToRecipeID   int64
}

type MeasurementUnit struct {
	Name        string
	Description pgtype.Text
}

type MeasurementUnitsTag struct {
	Measurement string
	Tag         string
}

type Recipe struct {
	ID                 int64
	AuthorID           int64
	ForkedFrom         pgtype.Int8
	Slug               string
	Description        pgtype.Text
	InitialPublishDate pgtype.Timestamp
}

type RecipeComment struct {
	ID       int64
	RecipeID int64
	AuthorID int64
	Content  string
	PostDate pgtype.Timestamp
}

type RecipeIngredient struct {
	ID         int64
	RevisionID int64
	Ingredient string
	Quantity   float32
	Unit       string
	Comment    pgtype.Text
}

type RecipeRevision struct {
	ID          int64
	RecipeID    int64
	ParentID    int64
	ChildID     pgtype.Int8
	Description pgtype.Text
	PublishDate pgtype.Timestamp
}

type RecipeStep struct {
	ID         int64
	RevisionID int64
	Content    string
	Index      int32
}

type Tag struct {
	Name        string
	Description pgtype.Text
}

type User struct {
	ID       int64
	Username string
	Email    string
	JoinDate pgtype.Timestamp
}
