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
