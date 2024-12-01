-- name: GetRecipeById :one
SELECT
  *
FROM
  recipes
WHERE
  recipes.id = $1
LIMIT 1;

-- name: GetRecipeBySlug :one
SELECT
  *
FROM
  recipes
WHERE
  recipes.slug = $1
LIMIT 1;

-- name: GetRecipesByAuthor :many
SELECT
  *
FROM
  recipes
WHERE
  recipes.author_id = $1
  AND recipes.id > $2
ORDER BY
  recipes.id
LIMIT $3;

-- name: GetRecipes :many
SELECT
  *
FROM
  recipes
WHERE
  recipes.id > $1
ORDER BY
  recipes.id
LIMIT $2;

-- name: CreateRecipe :one
INSERT INTO recipes (
  slug,
  author_id,
  description,
  forked_from
) VALUES (
	$1,
	$2,
  $3,
  $4
)
RETURNING *;
