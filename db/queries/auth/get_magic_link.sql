-- name: GetMagicLink :one
SELECT
  magic_links.id,
  magic_links.token,
  magic_links.user_id,
  magic_links.expiry
FROM
  magic_links
WHERE
  magic_links.id = $1 AND magic_links.token = $2;
