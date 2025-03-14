// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: upsert_measurement_unit.sql

package db

import (
	"context"
)

const upsertMeasurementUnit = `-- name: UpsertMeasurementUnit :one
WITH upsert AS (
  INSERT INTO
    measurement_units (
      name
    )
  VALUES (
    $1
  )
  ON CONFLICT (name)
  DO NOTHING
  RETURNING
    measurement_units.id,
    measurement_units.name
)
SELECT
  upsert.id,
	upsert.name
FROM
  upsert
UNION
SELECT
    measurement_units.id,
    measurement_units.name
FROM
  measurement_units
WHERE
  measurement_units.name = $1
`

type UpsertMeasurementUnitRow struct {
	ID   int64
	Name string
}

func (q *Queries) UpsertMeasurementUnit(ctx context.Context, name string) (UpsertMeasurementUnitRow, error) {
	row := q.db.QueryRow(ctx, upsertMeasurementUnit, name)
	var i UpsertMeasurementUnitRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
