-- name: GetMeasurementUnitFromIngredientId :one
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
LIMIT 1;
