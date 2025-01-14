-- name: GetIngredientFromRecipeIngredientId :one
SELECT
  ingredients.id,
  ingredients.name,
  ingredients.description
FROM
  recipe_ingredients
JOIN
  ingredients ON recipe_ingredients.ingredient = ingredients.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1;
