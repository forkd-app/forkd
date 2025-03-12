-- name: GetLatestRecipeRevisionByRecipeId :one
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date,
  photo
FROM
  recipe_revisions
WHERE
  recipe_id = $1
ORDER BY
  publish_date DESC
LIMIT 1;
