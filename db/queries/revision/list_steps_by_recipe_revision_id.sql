-- name: ListStepsByRecipeRevisionID :many
SELECT
    recipe_steps.id,
    recipe_steps.revision_id,
    recipe_steps.content,
    recipe_steps.index,
    recipe_steps.photo
FROM
    recipe_revisions
INNER JOIN
    recipe_steps ON recipe_revisions.id = recipe_steps.revision_id
WHERE
    recipe_revisions.id = sqlc.arg('revision_id')
ORDER BY
    recipe_steps.index;
