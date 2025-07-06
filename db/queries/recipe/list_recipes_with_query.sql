-- name: ListRecipesWithQuery :many
WITH rankings AS (
  SELECT
    r.*,
    ts_rank(setweight(to_tsvector('english', title), 'A') || setweight(to_tsvector('english', coalesce(recipe_description, '')), 'B'), websearch_to_tsquery('english', coalesce(sqlc.narg('query')))) AS rank
  FROM
    recipes r
  JOIN LATERAL (
    SELECT *
    FROM recipe_revisions
    WHERE
      -- If featured_revision is set, fetch that
      (r.featured_revision IS NOT NULL AND id = r.featured_revision)
      -- Otherwise, fetch the latest revision for this recipe
      OR (r.featured_revision IS NULL AND recipe_id = r.id)
    ORDER BY
      -- If featured_revision is set, this will be 1 row, so order doesn't matter
      -- If not, order by publish_date DESC to get the latest
      publish_date DESC
    LIMIT 1
  ) rev ON TRUE
)
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision,
  rank
FROM
  rankings
WHERE
  rank > 0
  AND
  CASE
    WHEN sqlc.narg('author_id')::uuid IS NOT NULL THEN author_id = sqlc.narg('author_id')::uuid
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('forked_from')::uuid IS NOT NULL THEN forked_from = sqlc.narg('forked_from')::uuid
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('current_user')::uuid IS NOT NULL AND sqlc.narg('private')::bool IS NOT NULL AND sqlc.narg('private')::bool THEN author_id = sqlc.narg('current_user')::uuid AND private = true
    ELSE private = false OR private IS NULL
  END
  AND
  CASE
    WHEN sqlc.narg('publish_start')::timestamp IS NOT NULL THEN initial_publish_date >= sqlc.narg('publish_start')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.narg('publish_end')::timestamp IS NOT NULL THEN initial_publish_date <= sqlc.narg('publish_end')::timestamp
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp > initial_publish_date
    ELSE true
  END
  AND
  CASE
    WHEN NOT sqlc.arg('sort_dir')::bool AND sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL THEN sqlc.narg('publish_cursor')::timestamp < initial_publish_date
    ELSE true
  END
  AND
  CASE
    WHEN sqlc.arg('sort_col')::text = 'slug' AND sqlc.arg('sort_dir')::bool AND sqlc.narg('slug_cursor')::text IS NOT NULL THEN sqlc.narg('slug_cursor')::text > slug
    ELSE true
  END
  AND
  CASE
    WHEN NOT sqlc.arg('sort_dir')::bool AND sqlc.arg('sort_col')::text = 'slug' AND sqlc.narg('slug_cursor')::text IS NOT NULL THEN sqlc.narg('slug_cursor')::text < slug
    ELSE true
  END
ORDER BY
  rank desc,
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND sqlc.arg('sort_dir')::bool THEN initial_publish_date END DESC,
  CASE WHEN sqlc.arg('sort_col')::text = 'publish_date' AND NOT sqlc.arg('sort_dir')::bool THEN initial_publish_date END ASC,
  CASE WHEN sqlc.arg('sort_col')::text = 'slug' AND sqlc.arg('sort_dir')::bool THEN slug END DESC,
  CASE WHEN sqlc.arg('sort_col')::text = 'slug' AND NOT sqlc.arg('sort_dir')::bool THEN slug END ASC
LIMIT sqlc.arg('limit');
