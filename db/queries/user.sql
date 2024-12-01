-- name: GetUserById :one
SELECT * FROM users WHERE users.id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE users.email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
	username,
	email
) VALUES (
	$1,
	$2
)
RETURNING *;
