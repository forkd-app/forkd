--Used when there is not a featured recipe
-- name: ListRecipesWithQuery :many
WITH LatesDateRecipe AS 
 (     
 SELECT 
    ROW_NUMBER() OVER (PARTITION BY recipe_id ORDER BY publish_date DESC) AS row_num,
    recipe_id,
    publish_date,
    recipe_description,
    title,
    id
FROM recipe_revisions

), AvailableIds as
(
Select LDR.recipe_description,
    LDR.title,
    LDR.id as revisionid,
    r.id as recipeid
    from recipes r
JOIN LatesDateRecipe LDR
ON r.id = LDR.recipe_id
Where LDR.row_num = 1
--and recipe_id = 'ae1f8b91-3659-4f7e-b484-dd54e7f3d2b3'
and r.featured_revision is null 

UNION  
--Used when there is a featured recipe
Select
    recipe_revisions.recipe_description,
    recipe_revisions.title,
    recipe_revisions.id as revisionid,
    r.id as recipeid
    from recipes r
JOIN recipe_revisions
    on r.featured_revision = recipe_revisions.id
)
,rankings AS (
  SELECT
    r.*,
    ts_rank(setweight(to_tsvector('english', title), 'A') || setweight(to_tsvector('english', coalesce(recipe_description, '')), 'B'), websearch_to_tsquery('english', sqlc.arg('query'))) AS rank
  FROM
    recipes r
  JOIN AvailableIds
  on r.id = AvailableIds.recipeid
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
