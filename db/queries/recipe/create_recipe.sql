-- name: CreateRecipe :one
INSERT INTO recipes (
  author_id,
  forked_from,
  slug,
  private
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision;
