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
-- name: ListRecipesByAuthor :many
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
  author_id = $1 AND id > $2
ORDER BY id
LIMIT $3;
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
  id > $1
ORDER BY id
LIMIT $2;
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