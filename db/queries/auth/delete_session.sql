-- name: DeleteSession :exec
DELETE FROM
  sessions
WHERE sessions.id = $1;
