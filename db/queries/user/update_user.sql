-- name: UpdateUser :one
UPDATE users
SET
    display_name = sqlc.arg('display_name'),
    email = sqlc.arg('email'),
    photo = sqlc.arg('photo')
WHERE id = sqlc.arg('id')
RETURNING
    id,
    display_name,
    email,
    join_date,
    updated_at,
    photo;
