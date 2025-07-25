-- name: GetUserById :one
SELECT
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at,
    users.photo
FROM users
WHERE users.id = sqlc.arg('id')
LIMIT 1;
