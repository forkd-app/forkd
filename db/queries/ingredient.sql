-- name: GetIngredientFromRecipeIngredientId :one
SELECT
  ingredients.id,
  ingredients.name,
  ingredients.description
FROM
  recipe_ingredients
JOIN
  ingredients ON recipe_ingredients.ingredient_id = ingredients.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1;

-- name: UpsertIngredient :one
WITH upsert AS (
  INSERT INTO
    ingredients (
      name
    )
  VALUES (
    $1
  )
  ON CONFLICT (name)
  DO NOTHING
  RETURNING
    ingredients.id,
    ingredients.name
)
SELECT
  upsert.id,
	upsert.name
FROM
  upsert
UNION
SELECT
    ingredients.id,
    ingredients.name
FROM
  ingredients
WHERE
  ingredients.name = $1;
