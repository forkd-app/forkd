-- name: CreateRecipeStep :one
INSERT INTO
recipe_steps (
    revision_id,
    content,
    index,
    photo
)
VALUES (
    sqlc.arg('revision_id'),
    sqlc.arg('content'),
    sqlc.arg('index'),
    sqlc.arg('photo')
)
RETURNING
    id,
    revision_id,
    content,
    index,
    photo;
