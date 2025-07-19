-- name: CreateUser :one
INSERT INTO
users (
    email,
    display_name
)
VALUES (
    sqlc.arg('email'),
    sqlc.arg('display_name')
)
RETURNING
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at,
    users.photo;
