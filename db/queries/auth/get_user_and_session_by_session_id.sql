-- name: GetUserBySessionId :one
SELECT
    sqlc.embed(users) AS user, -- noqa: RF02,RF04
    sqlc.embed(sessions) AS session -- noqa: RF02,RF04
FROM
    sessions
INNER JOIN
    users ON sessions.user_id = users.id
WHERE
    sessions.id = sqlc.arg('session_id')
LIMIT 1;
