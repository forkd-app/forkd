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
