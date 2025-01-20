-- name: GetUserById :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.email = $1 LIMIT 1;

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
    users.updated_at
)
SELECT
  upsert.id,
	upsert.display_name,
	upsert.email,
	upsert.join_date,
	upsert.updated_at
FROM
  upsert
UNION
SELECT
  users.id,
	users.display_name,
	users.email,
	users.join_date,
	users.updated_at
FROM
  users
WHERE
  users.email = $1;

-- name: GetUserBySessionId :one
SELECT
  sqlc.embed(users),
  sqlc.embed(sessions)
FROM
  sessions
JOIN
  users ON users.id = sessions.user_id
WHERE
  sessions.id = $1
LIMIT 1;

-- name: CreateSession :one
WITH sesh AS (
  INSERT INTO
    sessions (
      user_id,
      expiry
    )
  VALUES (
    $1,
    $2
  )
  RETURNING
    sessions.id,
    sessions.user_id
)
SELECT sqlc.embed(users), sesh.id FROM sesh INNER JOIN users ON sesh.user_id = users.id;

-- name: DeleteSession :exec
DELETE FROM
  sessions
WHERE sessions.id = $1;

-- name: CreateMagicLink :one
INSERT INTO
  magic_links (
    user_id,
    token,
    expiry
  )
VALUES (
  $1,
  $2,
  $3
)
RETURNING
  magic_links.id,
  magic_links.token;

-- name: GetMagicLink :one
SELECT
  magic_links.id,
  magic_links.token,
  magic_links.user_id,
  magic_links.expiry
FROM
  magic_links
WHERE
  magic_links.id = $1 AND magic_links.token = $2
LIMIT 1;
