-- name: SeedUser :one
INSERT INTO
  users (
    email,
    display_name,
    join_date
  )
VALUES (
  $1,
  $2,
  $3
)
RETURNING
  users.id,
  users.display_name,
  users.email,
  users.join_date,
  users.updated_at,
  users.photo;
