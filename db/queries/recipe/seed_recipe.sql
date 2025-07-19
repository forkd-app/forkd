-- name: SeedRecipe :one
INSERT INTO recipes (
    author_id,
    forked_from,
    slug,
    private,
    initial_publish_date
) VALUES (
    sqlc.arg('author_id'),
    sqlc.arg('forked_from'),
    sqlc.arg('slug'),
    sqlc.arg('private'),
    sqlc.arg('initial_publish_date')
) RETURNING
    id,
    author_id,
    slug,
    private,
    initial_publish_date,
    forked_from,
    featured_revision;
