-- name: ListRecipesWithQuery :many
--RevisionWithRowNumbers takes the recipe revision and add row numbers to be used later
--This change is used to help us choose the most recent revision in the absence of a featured revision
WITH RevisionsWithRowNumbers
AS (
	SELECT ROW_NUMBER() OVER (
			PARTITION BY recipe_id ORDER BY publish_date DESC
			) AS row_num
		,recipe_id
		,recipe_description
		,title
		,id
	FROM recipe_revisions
	)
	--LatestRevision takes the RowNumbers partitioned above. Row_Num 1 is the most recent based on revisions ordered by date.
	,LatestRevision
AS (
	SELECT LDR.recipe_description
		,LDR.title
		,LDR.id AS revisionid
	FROM recipes r
	JOIN RevisionsWithRowNumbers LDR ON r.id = LDR.recipe_id
	WHERE LDR.row_num = 1
		AND r.featured_revision IS NULL
	
	UNION
	
	--Used when there is a featured recipe
	SELECT recipe_revisions.recipe_description
		,recipe_revisions.title
		,recipe_revisions.id AS revisionid
	FROM recipes r
	JOIN recipe_revisions ON r.featured_revision = recipe_revisions.id
	)
	,rankings
AS (
	SELECT r.*
		,ts_rank(setweight(to_tsvector('english', title), 'A') || setweight(to_tsvector('english', coalesce(recipe_description, '')), 'B'), websearch_to_tsquery('english', sqlc.arg('query'))) AS rank
	FROM recipes r
	JOIN LatestRevision ON r.id = LatestRevision.recipeid
	)
SELECT id
	,author_id
	,slug
	,private
	,initial_publish_date
	,forked_from
	,featured_revision
	,rank
FROM rankings
WHERE rank > 0
	AND CASE 
		WHEN sqlc.narg('author_id')::uuid IS NOT NULL
			THEN author_id = sqlc.narg('author_id')::uuid
		ELSE true
		END
	AND CASE 
		WHEN sqlc.narg('forked_from')::uuid IS NOT NULL
			THEN forked_from = sqlc.narg('forked_from')::uuid
		ELSE true
		END
	AND CASE 
		WHEN sqlc.narg('current_user')::uuid IS NOT NULL
			AND sqlc.narg('private')::bool IS NOT NULL
			AND sqlc.narg('private')::bool
			THEN author_id = sqlc.narg('current_user')::uuid
				AND private = true
		ELSE private = false
			OR private IS NULL
		END
	AND CASE 
		WHEN sqlc.narg('publish_start')::TIMESTAMP IS NOT NULL
			THEN initial_publish_date >= sqlc.narg('publish_start')::TIMESTAMP
		ELSE true
		END
	AND CASE 
		WHEN sqlc.narg('publish_end')::TIMESTAMP IS NOT NULL
			THEN initial_publish_date <= sqlc.narg('publish_end')::TIMESTAMP
		ELSE true
		END
	AND CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'publish_date'
			AND sqlc.arg('sort_dir')::bool
			AND sqlc.narg('publish_cursor')::TIMESTAMP IS NOT NULL
			THEN sqlc.narg('publish_cursor')::TIMESTAMP > initial_publish_date
		ELSE true
		END
	AND CASE 
		WHEN NOT sqlc.arg('sort_dir')::bool
			AND sqlc.arg('sort_col')::TEXT = 'publish_date'
			AND sqlc.narg('publish_cursor')::TIMESTAMP IS NOT NULL
			THEN sqlc.narg('publish_cursor')::TIMESTAMP < initial_publish_date
		ELSE true
		END
	AND CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'slug'
			AND sqlc.arg('sort_dir')::bool
			AND sqlc.narg('slug_cursor')::TEXT IS NOT NULL
			THEN sqlc.narg('slug_cursor')::TEXT > slug
		ELSE true
		END
	AND CASE 
		WHEN NOT sqlc.arg('sort_dir')::bool
			AND sqlc.arg('sort_col')::TEXT = 'slug'
			AND sqlc.narg('slug_cursor')::TEXT IS NOT NULL
			THEN sqlc.narg('slug_cursor')::TEXT < slug
		ELSE true
		END
ORDER BY rank DESC
	,CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'publish_date'
			AND sqlc.arg('sort_dir')::bool
			THEN initial_publish_date
		END DESC
	,CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'publish_date'
			AND NOT sqlc.arg('sort_dir')::bool
			THEN initial_publish_date
		END ASC
	,CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'slug'
			AND sqlc.arg('sort_dir')::bool
			THEN slug
		END DESC
	,CASE 
		WHEN sqlc.arg('sort_col')::TEXT = 'slug'
			AND NOT sqlc.arg('sort_dir')::bool
			THEN slug
		END ASC LIMIT sqlc.arg('limit');

