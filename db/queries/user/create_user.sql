-- name: CreateUser :one
INSERT INTO
  users (
    email,
    display_name
  )
VALUES (
  $1,
  $2
)
RETURNING
  users.id,
  users.display_name,
  users.email,
  users.join_date,
  users.updated_at,
  users.photo;
