// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: get_ingredient_by_id.sql

package db

import (
	"context"
)

const getIngredientById = `-- name: GetIngredientById :one
SELECT
  ingredients.id,
  ingredients.name,
  ingredients.description
FROM
  ingredients
WHERE
  ingredients.id = $1
LIMIT 1
`

func (q *Queries) GetIngredientById(ctx context.Context, id int64) (Ingredient, error) {
	row := q.db.QueryRow(ctx, getIngredientById, id)
	var i Ingredient
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
