-- name: CreateRevision :one
INSERT INTO recipe_revisions (
    recipe_id,
    parent_id,
    recipe_description,
    change_comment,
    title,
    photo
)
VALUES (
    sqlc.arg('recipe_id'),
    sqlc.arg('parent_id'),
    sqlc.arg('recipe_description'),
    sqlc.arg('change_comment'),
    sqlc.arg('title'),
    sqlc.arg('photo')
)
RETURNING
    id,
    recipe_id,
    parent_id,
    recipe_description,
    change_comment,
    title,
    publish_date,
    photo;
