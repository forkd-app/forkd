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
    id = sqlc.arg('id')
LIMIT 1;
