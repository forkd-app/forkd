-- name: GetRecipeBySlug :one
SELECT
  recipes.id,
  recipes.author_id,
  recipes.slug,
  recipes.private,
  recipes.initial_publish_date,
  recipes.forked_from,
  recipes.featured_revision
FROM
  recipes
JOIN
  users ON users.id = recipes.author_id
WHERE
  lower(recipes.slug) = lower(sqlc.arg(slug)::text)
  AND
  lower(users.display_name) = lower(sqlc.arg(display_name)::text)
LIMIT 1;
