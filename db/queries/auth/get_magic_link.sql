-- name: GetMagicLink :one
SELECT
    magic_links.id,
    magic_links.token,
    magic_links.user_id,
    magic_links.expiry
FROM
    magic_links
WHERE
    magic_links.id = sqlc.arg('id') AND magic_links.token = sqlc.arg('token');
