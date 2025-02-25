-- name: ListIngredientsByRecipeRevisionID :many
SELECT
  recipe_ingredients.id,
  recipe_ingredients.revision_id,
  recipe_ingredients.ingredient_id,
  recipe_ingredients.quantity,
  recipe_ingredients.measurement_unit_id,
  recipe_ingredients.comment
FROM
  recipe_revisions
JOIN
  recipe_ingredients ON recipe_revisions.id = recipe_ingredients.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY recipe_ingredients.id;
