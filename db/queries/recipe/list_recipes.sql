-- name: ListRecipes :many
SELECT
    id,
    author_id,
    slug,
    private,
    initial_publish_date,
    forked_from,
    featured_revision
FROM
    recipes
WHERE
    CASE
        WHEN
            sqlc.narg('author_id')::uuid IS NOT NULL
            THEN author_id = sqlc.narg('author_id')::uuid
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            sqlc.narg('forked_from')::uuid IS NOT NULL
            THEN forked_from = sqlc.narg('forked_from')::uuid
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            sqlc.narg('current_user')::uuid IS NOT NULL
            AND sqlc.narg('private')::bool IS NOT NULL
            AND sqlc.narg('private')::bool
            THEN author_id = sqlc.narg('current_user')::uuid AND private = TRUE
        ELSE private = FALSE OR private IS NULL
    END
    AND
    CASE
        WHEN
            sqlc.narg('publish_start')::timestamp IS NOT NULL
            THEN initial_publish_date >= sqlc.narg('publish_start')::timestamp
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            sqlc.narg('publish_end')::timestamp IS NOT NULL
            THEN initial_publish_date <= sqlc.narg('publish_end')::timestamp
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.arg('sort_dir')::bool
            AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL
            THEN sqlc.narg('publish_cursor')::timestamp > initial_publish_date
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            NOT sqlc.arg('sort_dir')::bool
            AND sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.narg('publish_cursor')::timestamp IS NOT NULL
            THEN sqlc.narg('publish_cursor')::timestamp < initial_publish_date
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug'
            AND sqlc.arg('sort_dir')::bool
            AND sqlc.narg('slug_cursor')::text IS NOT NULL
            THEN sqlc.narg('slug_cursor')::text > slug
        ELSE TRUE
    END
    AND
    CASE
        WHEN
            NOT sqlc.arg('sort_dir')::bool
            AND sqlc.arg('sort_col')::text = 'slug'
            AND sqlc.narg('slug_cursor')::text IS NOT NULL
            THEN sqlc.narg('slug_cursor')::text < slug
        ELSE TRUE
    END
ORDER BY
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND sqlc.arg('sort_dir')::bool
            THEN initial_publish_date
    END DESC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'publish_date'
            AND NOT sqlc.arg('sort_dir')::bool
            THEN initial_publish_date
    END ASC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug' AND sqlc.arg('sort_dir')::bool
            THEN slug
    END DESC,
    CASE
        WHEN
            sqlc.arg('sort_col')::text = 'slug'
            AND NOT sqlc.arg('sort_dir')::bool
            THEN slug
    END ASC
LIMIT sqlc.arg('limit');
