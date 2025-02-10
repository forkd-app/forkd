-- name: GetRecipeById :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  id = $1
LIMIT 1;
-- name: GetRecipeBySlug :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  slug = $1
LIMIT 1;
-- name: ListRecipes :many
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  CASE
    WHEN sqlc.narg('author_id')::uuid IS NOT NULL THEN author_id = sqlc.narg('author_id')::uuid
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('forked_from')::uuid IS NOT NULL THEN forked_from = sqlc.narg('forked_from')::uuid
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('private')::bool IS NOT NULL THEN private = sqlc.narg('private')::bool
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('publish_start')::timestamp IS NOT NULL THEN initial_publish_date > sqlc.narg('publish_start')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('publish_end')::timestamp IS NOT NULL THEN initial_publish_date < sqlc.narg('publish_end')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp > initial_publish_date
    ELSE true
  END
  AND
  CASE
    WHEN NOT sqlc.arg('sort_dir')::bool AND sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp < initial_publish_date
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.arg('sort_col')::text = 'slug' AND sqlc.arg('sort_dir')::bool AND sqlc.narg('slug_cursor')::text IS NOT NULL THEN sqlc.narg('slug_cursor')::text > slug
    ELSE true
  END
  AND
  CASE
    WHEN NOT sqlc.arg('sort_dir')::bool AND sqlc.arg('sort_col')::text = 'slug' AND sqlc.narg('slug_cursor')::text IS NOT NULL THEN sqlc.narg('slug_cursor')::text < slug
    ELSE true
  END
ORDER BY
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool THEN initial_publish_date END DESC,
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND NOT sqlc.arg('sort_dir')::bool THEN initial_publish_date END ASC,
  CASE WHEN sqlc.arg('sort_col')::text = 'slug' AND sqlc.arg('sort_dir')::bool THEN slug END DESC,
  CASE WHEN sqlc.arg('sort_col')::text = 'slug' AND NOT sqlc.arg('sort_dir')::bool THEN slug END ASC
LIMIT sqlc.arg('limit');
-- name: CreateRecipe :one
INSERT INTO recipes (
  author_id,
  forked_from,
  slug,
  private
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision;
-- name: GetRecipeByRevisionID :one
SELECT
  recipes.id,
  recipes.author_id,
  recipes.slug,
  recipes.private,
  recipes.initial_publish_date,
  recipes.forked_from,
  recipes.featured_revision
FROM
  recipe_revisions
JOIN
  recipes ON recipe_revisions.recipe_id = recipes.id
WHERE
  recipes.id = $1
LIMIT 1;
-- name: GetRecipeRevisionByParentID :one
SELECT
  parent.id,
  parent.recipe_id,
  parent.parent_id,
  parent.recipe_description,
  parent.change_comment,
  parent.title,
  parent.publish_date
FROM
  recipe_revisions child
JOIN
  recipe_revisions parent ON child.parent_id = parent.id
WHERE
  recipe_revisions.id = $1
LIMIT 1;
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
-- name: ListStepsByRecipeRevisionID :many
SELECT
  recipe_steps.id,
  recipe_steps.revision_id,
  recipe_steps.content,
  recipe_steps.index
FROM
  recipe_revisions
JOIN
  recipe_steps ON recipe_revisions.id = recipe_steps.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY
  recipe_steps.index;
-- name: UpdateRecipe :one
UPDATE recipes
SET
  slug = coalesce($1, slug),
  private = coalesce($2, private),
  featured_revision = coalesce($3, featured_revision)
WHERE id = $4
RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision;
