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
