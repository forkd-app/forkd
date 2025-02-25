-- name: GetAuthorByRecipeId :one
SELECT
  users.id,
  users.display_name,
  users.email,
  users.join_date,
  users.updated_at,
  users.photo
FROM
  users
JOIN recipes ON users.id = recipes.author_id
WHERE
  recipes.id = $1
LIMIT 1;

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
