-- name: DeleteMagicLinkById :exec
DELETE FROM
magic_links
WHERE magic_links.id = sqlc.arg('id');
