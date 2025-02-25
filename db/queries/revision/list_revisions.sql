-- name: ListRevisions :many
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date,
  photo
FROM
  recipe_revisions
WHERE
  CASE
    WHEN sqlc.narg('recipe_id')::uuid IS NOT NULL THEN sqlc.narg('recipe_id')::uuid = recipe_id
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('parent_id')::uuid IS NOT NULL THEN sqlc.narg('parent_id')::uuid = parent_id
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('publish_start')::timestamp IS NOT NULL THEN publish_date >= sqlc.narg('publish_start')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('publish_end')::timestamp IS NOT NULL THEN publish_date <= sqlc.narg('publish_end')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp > publish_date
    ELSE true
  END
  AND
  CASE
    WHEN NOT sqlc.arg('sort_dir')::bool AND sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp < publish_date
    ELSE true
  END
ORDER BY
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool THEN publish_date END DESC,
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND NOT sqlc.arg('sort_dir')::bool THEN publish_date END ASC
LIMIT sqlc.arg('limit'); -- Limit for pagination
