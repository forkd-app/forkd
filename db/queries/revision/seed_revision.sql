-- name: SeedRevision :one
INSERT INTO recipe_revisions (
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  photo,
  publish_date
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
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
