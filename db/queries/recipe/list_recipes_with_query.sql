-- name: ListRecipesWithQuery :many
--RevisionWithRowNumbers takes the recipe revision 
--and add row numbers to be used later
--This change is used to help us choose the most recent revision 
--in the absence of a featured revision
WITH REVISIONSWITHROWNUMBERS AS (
    SELECT
        RECIPE_ID,
        RECIPE_DESCRIPTION,
        TITLE,
        ID,
        ROW_NUMBER() OVER (
            PARTITION BY RECIPE_ID ORDER BY PUBLISH_DATE DESC
        ) AS ROW_NUM
    FROM RECIPE_REVISIONS
),

--LatestRevision takes the RowNumbers partitioned above.
--Row_Num 1 is the most recent based on revisions ordered by date.
LATESTREVISION AS (
    SELECT
        LDR.RECIPE_DESCRIPTION,
        LDR.TITLE,
        LDR.ID AS REVISIONID
    FROM RECIPES AS R
    INNER JOIN REVISIONSWITHROWNUMBERS AS LDR ON R.ID = LDR.RECIPE_ID
    WHERE
        LDR.ROW_NUM = 1
        AND R.FEATURED_REVISION IS NULL
    UNION
    --Used when there is a featured recipe
    SELECT
        RECIPE_REVISIONS.RECIPE_DESCRIPTION,
        RECIPE_REVISIONS.TITLE,
        RECIPE_REVISIONS.ID AS REVISIONID
    FROM RECIPES AS R
    INNER JOIN RECIPE_REVISIONS ON R.FEATURED_REVISION = RECIPE_REVISIONS.ID
),

RANKINGS AS (
    SELECT
        R.*,
        TS_RANK(
            SETWEIGHT(TO_TSVECTOR('english', LATESTREVISION.TITLE), 'A')
            || SETWEIGHT(
                TO_TSVECTOR(
                    'english',
                    COALESCE(LATESTREVISION.RECIPE_DESCRIPTION, '')
                ),
                'B'
            ),
            WEBSEARCH_TO_TSQUERY('english', sqlc.arg('query'))
        ) AS RANK
    FROM RECIPES AS R
    INNER JOIN LATESTREVISION ON R.ID = LATESTREVISION.RECIPEID
)

SELECT
    ID,
    AUTHOR_ID,
    SLUG,
    PRIVATE,
    INITIAL_PUBLISH_DATE,
    FORKED_FROM,
    FEATURED_REVISION,
    RANK
FROM RANKINGS
WHERE
    RANK > 0
    AND CASE
        WHEN sqlc.narg('author_id')::uuid IS NOT NULL
            THEN AUTHOR_ID = sqlc.narg('author_id')::uuid
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.narg('forked_from')::uuid IS NOT NULL
            THEN FORKED_FROM = sqlc.narg('forked_from')::uuid
        ELSE TRUE
    END
    AND CASE
        WHEN
            sqlc.narg('current_user')::uuid IS NOT NULL
            AND sqlc.narg('private')::bool IS NOT NULL
            AND sqlc.narg('private')::bool
            THEN
                AUTHOR_ID = sqlc.narg('current_user')::uuid
                AND PRIVATE = TRUE
        ELSE
            PRIVATE = FALSE
            OR PRIVATE IS NULL
    END
    AND CASE
        WHEN sqlc.narg('publish_start')::timestamp IS NOT NULL
            THEN INITIAL_PUBLISH_DATE >= sqlc.narg('publish_start')::timestamp
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.narg('publish_end')::timestamp IS NOT NULL
            THEN INITIAL_PUBLISH_DATE <= sqlc.narg('publish_end')::timestamp
        ELSE TRUE
    END
    AND CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.arg('sort_dir')::bool
            AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL
            THEN
                sqlc.narg('publish_cursor')::timestamp > INITIAL_PUBLISH_DATE
        ELSE TRUE
    END
    AND CASE
        WHEN
            NOT sqlc.arg('sort_dir')::bool
            AND sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL
            THEN
                sqlc.narg('publish_cursor')::timestamp < INITIAL_PUBLISH_DATE
        ELSE TRUE
    END
    AND CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug'
            AND sqlc.arg('sort_dir')::bool
            AND sqlc.narg('slug_cursor')::text IS NOT NULL
            THEN sqlc.narg('slug_cursor')::text > SLUG
        ELSE TRUE
    END
    AND CASE
        WHEN
            NOT sqlc.arg('sort_dir')::bool
            AND sqlc.arg('sort_col')::text = 'slug'
            AND sqlc.narg('slug_cursor')::text IS NOT NULL
            THEN sqlc.narg('slug_cursor')::text < SLUG
        ELSE TRUE
    END
ORDER BY
    RANK DESC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.arg('sort_dir')::bool
            THEN INITIAL_PUBLISH_DATE
    END DESC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND NOT sqlc.arg('sort_dir')::bool
            THEN INITIAL_PUBLISH_DATE
    END ASC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug'
            AND sqlc.arg('sort_dir')::bool
            THEN SLUG
    END DESC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug'
            AND NOT sqlc.arg('sort_dir')::bool
            THEN SLUG
    END ASC
LIMIT sqlc.arg('limit');
