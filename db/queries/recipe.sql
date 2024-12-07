-- name: GetRecipeById :one
SELECT
  recipes.id,
  recipes.slug,
  recipes.description,
  recipes.author_id,
  recipes.forked_from,
  recipes.initial_publish_date
FROM
  recipes
WHERE
  recipes.id = $1
LIMIT 1;
-- name: GetRecipeBySlug :one
SELECT
recipes.id,
recipes.slug,
recipes.description,
recipes.author_id,
recipes.forked_from,
recipes.initial_publish_date
FROM
recipes
WHERE
recipes.slug = $1
LIMIT 1;
-- name: GetRecipesByAuthor :many
SELECT
recipes.id,
recipes.slug,
recipes.description,
recipes.author_id,
recipes.forked_from,
recipes.initial_publish_date
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
INSERT INTO recipes(
slug,
author_id,
description,
forked_from
)
VALUES(
$1,
$2,
$3,
$4
) RETURNING
recipes.id,
recipes.slug,
recipes.description,
recipes.author_id,
recipes.forked_from,
recipes.initial_publish_date;
