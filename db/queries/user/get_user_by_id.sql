-- name: GetUserById :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at, users.photo FROM users WHERE users.id = $1 LIMIT 1;
