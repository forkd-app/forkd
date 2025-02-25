-- name: UpsertUser :one
WITH upsert AS (
  INSERT INTO
    users (
      email,
      display_name
    )
  VALUES (
    $1,
    $2
  )
  ON CONFLICT (email)
  DO NOTHING
  RETURNING
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at,
    users.photo
)
SELECT
  upsert.id,
	upsert.display_name,
	upsert.email,
	upsert.join_date,
	upsert.updated_at,
  upsert.photo
FROM
  upsert
UNION
SELECT
  users.id,
	users.display_name,
	users.email,
	users.join_date,
	users.updated_at,
  users.photo
FROM
  users
WHERE
  users.email = $1;
