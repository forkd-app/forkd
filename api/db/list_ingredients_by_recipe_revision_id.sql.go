// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: list_ingredients_by_recipe_revision_id.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const listIngredientsByRecipeRevisionID = `-- name: ListIngredientsByRecipeRevisionID :many
SELECT
  recipe_ingredients.id,
  recipe_ingredients.revision_id,
  recipe_ingredients.ingredient_id,
  recipe_ingredients.quantity,
  recipe_ingredients.measurement_unit_id,
  recipe_ingredients.comment
FROM
  recipe_revisions
JOIN
  recipe_ingredients ON recipe_revisions.id = recipe_ingredients.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY recipe_ingredients.id
`

func (q *Queries) ListIngredientsByRecipeRevisionID(ctx context.Context, revisionID pgtype.UUID) ([]RecipeIngredient, error) {
	rows, err := q.db.Query(ctx, listIngredientsByRecipeRevisionID, revisionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecipeIngredient
	for rows.Next() {
		var i RecipeIngredient
		if err := rows.Scan(
			&i.ID,
			&i.RevisionID,
			&i.IngredientID,
			&i.Quantity,
			&i.MeasurementUnitID,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
