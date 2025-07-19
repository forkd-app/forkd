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
INNER JOIN
    users ON recipes.author_id = users.id
WHERE
    LOWER(recipes.slug) = LOWER(sqlc.arg('slug')::text)
    AND
    LOWER(users.display_name) = LOWER(sqlc.arg('display_name')::text)
LIMIT 1;
