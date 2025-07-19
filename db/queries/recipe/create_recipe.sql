-- name: CreateRecipe :one
INSERT INTO recipes (
    author_id,
    forked_from,
    slug,
    private
) VALUES (
    sqlc.arg('author_id'),
    sqlc.arg('forked_from'),
    sqlc.arg('slug'),
    sqlc.arg('private')
) RETURNING
    id,
    author_id,
    slug,
    private,
    initial_publish_date,
    forked_from,
    featured_revision;
