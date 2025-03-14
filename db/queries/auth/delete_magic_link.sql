-- name: DeleteMagicLinkById :exec
DELETE FROM
  magic_links
WHERE magic_links.id = $1;
