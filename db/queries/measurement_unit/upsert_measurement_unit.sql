-- name: UpsertMeasurementUnit :one
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
  measurement_units.name = $1;
