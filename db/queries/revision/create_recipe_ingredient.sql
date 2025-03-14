-- name: CreateRecipeIngredient :one
INSERT INTO
  recipe_ingredients (
    revision_id,
    ingredient_id,
    measurement_unit_id,
    quantity,
    comment
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING
  recipe_ingredients.id,
  recipe_ingredients.revision_id,
  recipe_ingredients.ingredient_id,
  recipe_ingredients.quantity,
  recipe_ingredients.measurement_unit_id,
  recipe_ingredients.comment;
