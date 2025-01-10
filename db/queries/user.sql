-- name: GetUserById :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
	display_name,
	email
) VALUES (
	$1,
	$2
)
RETURNING users.id, users.display_name, users.email, users.join_date, users.updated_at;
