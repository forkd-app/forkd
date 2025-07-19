-- name: HasUserRatedRevision :one
SELECT EXISTS(
    SELECT star_value
    FROM ratings
    WHERE
        revision_id = sqlc.arg('revision_id')::uuid
        AND
        user_id = sqlc.arg('user_id')::uuid
);
