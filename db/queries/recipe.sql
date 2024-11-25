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
