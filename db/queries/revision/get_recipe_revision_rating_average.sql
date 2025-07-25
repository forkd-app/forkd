-- name: GetRecipeRevisionRatingAverage :one
SELECT AVG(star_value) AS average
FROM ratings
WHERE revision_id = sqlc.arg('revision_id')::uuid;
