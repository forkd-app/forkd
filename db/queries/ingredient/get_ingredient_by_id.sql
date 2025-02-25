-- name: GetIngredientById :one
SELECT
  ingredients.id,
  ingredients.name,
  ingredients.description
FROM
  ingredients
WHERE
  ingredients.id = $1
LIMIT 1;
