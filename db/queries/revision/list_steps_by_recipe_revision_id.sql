-- name: ListStepsByRecipeRevisionID :many
SELECT
  recipe_steps.id,
  recipe_steps.revision_id,
  recipe_steps.content,
  recipe_steps.index,
  recipe_steps.photo
FROM
  recipe_revisions
JOIN
  recipe_steps ON recipe_revisions.id = recipe_steps.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY
  recipe_steps.index;
