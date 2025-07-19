-- name: CreateSession :one
WITH sesh AS (
    INSERT INTO
    sessions (
        user_id,
        expiry
    )
    VALUES (
        sqlc.arg('user_id'),
        sqlc.arg('expiry')
    )
    RETURNING
        sessions.id,
        sessions.user_id
)

SELECT
    sesh.id,
    sqlc.embed(users) AS user -- noqa: RF02,RF04
FROM sesh
INNER JOIN users ON sesh.user_id = users.id;
