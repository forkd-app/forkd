-- name: GetRecipeRevisionById :one
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
    id = sqlc.arg('id')
LIMIT 1;
