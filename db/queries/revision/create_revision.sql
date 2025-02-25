-- name: CreateRevision :one
INSERT INTO recipe_revisions (
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  photo
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date,
  photo;
