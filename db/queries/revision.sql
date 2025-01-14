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
  recipe_id = $1
  AND id > $2 -- Cursor for pagination
ORDER BY id
LIMIT $3; -- Limit for pagination

