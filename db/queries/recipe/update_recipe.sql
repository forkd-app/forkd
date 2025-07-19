-- name: UpdateRecipe :one
UPDATE recipes
SET
    slug = coalesce(sqlc.arg('slug'), slug),
    private = coalesce(sqlc.arg('private'), private),
    featured_revision
    = coalesce(sqlc.arg('featured_revision'), featured_revision)
WHERE id = sqlc.arg('id')
RETURNING
    id,
    author_id,
    slug,
    private,
    initial_publish_date,
    forked_from,
    featured_revision;
