-- name: UpdateRecipe :one
UPDATE recipes
SET
  slug = coalesce($1, slug),
  private = coalesce($2, private),
  featured_revision = coalesce($3, featured_revision)
WHERE id = $4
RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision;
