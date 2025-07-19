-- name: UpsertRating :one
INSERT INTO ratings (
    revision_id,
    user_id,
    star_value
)
VALUES (
    sqlc.arg('revision_id')::uuid,
    sqlc.arg('user_id')::uuid,
    sqlc.arg('star_value')::smallint
)
ON CONFLICT (revision_id, user_id)
DO UPDATE SET star_value = excluded.star_value
RETURNING
    revision_id,
    user_id,
    star_value;
