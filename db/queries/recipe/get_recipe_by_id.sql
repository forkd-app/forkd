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
