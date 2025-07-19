-- name: DeleteSession :exec
DELETE FROM
sessions
WHERE sessions.id = sqlc.arg('id');
