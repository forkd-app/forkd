-- name: CreateRecipe :one
INSERT INTO recipes (
  author_id,
  forked_from,
  slug,
  private,
  initial_publish_date
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision;
