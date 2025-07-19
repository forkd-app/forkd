-- name: SeedUser :one
INSERT INTO
users (
    email,
    display_name,
    join_date
)
VALUES (
    sqlc.arg('email'),
    sqlc.arg('display_name'),
    sqlc.arg('join_date')
)
RETURNING
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at,
    users.photo;
