-- name: CreateRecipeStep :one
INSERT INTO
  recipe_steps (
    revision_id,
    content,
    index,
    photo
  )
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING
  id,
  revision_id,
  content,
  index,
  photo;
