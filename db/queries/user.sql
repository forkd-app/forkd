-- name: GetUserById :one
SELECT users.id, users.email, users.username, users.join_date FROM users WHERE users.id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT users.id, users.email, users.username, users.join_date FROM users WHERE users.email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
	username,
	email
) VALUES (
	$1,
	$2
)
RETURNING users.id, users.email, users.username, users.join_date;
