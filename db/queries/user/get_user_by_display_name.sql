-- name: GetUserByDisplayName :one
SELECT
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at,
    users.photo
FROM users
WHERE users.display_name = sqlc.arg('display_name')
LIMIT 1;
