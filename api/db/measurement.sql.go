// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: measurement.sql

package db

import (
	"context"
)

const getMeasurementUnitFromIngredientId = `-- name: GetMeasurementUnitFromIngredientId :one
SELECT
  measurement_units.id,
  measurement_units.name,
  measurement_units.description
FROM
  recipe_ingredients
JOIN
  measurement_units ON recipe_ingredients.unit = measurement_units.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1
`

func (q *Queries) GetMeasurementUnitFromIngredientId(ctx context.Context, id int64) (MeasurementUnit, error) {
	row := q.db.QueryRow(ctx, getMeasurementUnitFromIngredientId, id)
	var i MeasurementUnit
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
