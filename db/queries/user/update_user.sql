-- name: UpdateUser :one
UPDATE users
SET display_name = $2, email = $3, photo = $4
WHERE id = $1
RETURNING
  id,
  display_name,
  email,
  join_date,
  updated_at,
  photo;
