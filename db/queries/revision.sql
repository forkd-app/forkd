-- name: GetRecipeRevisionById :one
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date
FROM
  recipe_revisions
WHERE
  id = $1
LIMIT 1;
-- name: GetRecipeRevisionByStepId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipe_steps
JOIN
  recipe_revisions ON recipe_steps.revision_id = recipe_revisions.id
WHERE
  recipe_steps.id = $1
LIMIT 1;
-- name: GetRecipeRevisionByIngredientId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipe_ingredients
JOIN
  recipe_revisions ON recipe_ingredients.revision_id = recipe_revisions.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1;

-- name: ListRecipeRevisions :many
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date
FROM
  recipe_revisions
WHERE
  recipe_id = sqlc.arg('recipe_id')
  AND CASE
    WHEN sqlc.narg('id')::uuid IS NOT NULL THEN id > sqlc.narg('id')::uuid
    ELSE true
  END
ORDER BY id
LIMIT sqlc.arg('limit'); -- Limit for pagination

-- name: GetForkedFromRevisionByRecipeId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipes
JOIN recipe_revisions ON recipes.forked_from = recipe_revisions.id
WHERE
  recipes.id = $1
LIMIT 1;

-- name: GetFeaturedRevisionByRecipeId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipes
JOIN recipe_revisions ON recipes.featured_revision = recipe_revisions.id
WHERE
  recipes.id = $1
LIMIT 1;

-- name: CreateRevision :one
INSERT INTO recipe_revisions (
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date;

-- name: CreateRevisionIngredient :one
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

-- name: CreateRevisionStep :one
INSERT INTO
  recipe_steps (
    revision_id,
    content,
    index
  )
VALUES (
  $1,
  $2,
  $3
)
RETURNING
  recipe_steps.id,
  recipe_steps.revision_id,
  recipe_steps.content,
  recipe_steps.index;
