-- name: GetRecipeRevisionById :one
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
  id = $1
LIMIT 1;
