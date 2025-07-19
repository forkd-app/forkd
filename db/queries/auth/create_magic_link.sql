-- name: CreateMagicLink :one
INSERT INTO
magic_links (
    user_id,
    token,
    expiry
)
VALUES (
    sqlc.arg('user_id'),
    sqlc.arg('token'),
    sqlc.arg('expiry')
)
RETURNING
    magic_links.id,
    magic_links.token;
