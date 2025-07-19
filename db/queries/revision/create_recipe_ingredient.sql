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
    sqlc.arg('revision_id'),
    sqlc.arg('ingredient_id'),
    sqlc.arg('measurement_unit_id'),
    sqlc.arg('quantity'),
    sqlc.arg('comment')
)
RETURNING
    recipe_ingredients.id,
    recipe_ingredients.revision_id,
    recipe_ingredients.ingredient_id,
    recipe_ingredients.quantity,
    recipe_ingredients.measurement_unit_id,
    recipe_ingredients.comment;
