-- name: GetMeasurementUnitFromIngredientId :one
SELECT
  measurement_units.id,
  measurement_units.name,
  measurement_units.description
FROM
  recipe_ingredients
JOIN
  measurement_units ON recipe_ingredients.measurement_unit_id = measurement_units.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1;

-- name: UpsertMeasurement :one
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
