-- name: GetIngredientById :one
SELECT
    ingredients.id,
    ingredients.name,
    ingredients.description
FROM
    ingredients
WHERE
    ingredients.id = sqlc.arg('id')
LIMIT 1;
