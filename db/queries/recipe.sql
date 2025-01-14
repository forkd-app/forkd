-- name: GetRecipeById :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  id = $1
LIMIT 1;
-- name: GetRecipeBySlug :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  slug = $1
LIMIT 1;
-- name: ListRecipesByAuthor :many
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  author_id = $1 AND id > $2
ORDER BY id
LIMIT $3;
-- name: ListRecipes :many
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  id > $1
ORDER BY id
LIMIT $2;
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

-- name: GetRecipeWithAuthorById :one
SELECT
  r.id AS recipe_id,
  r.slug AS recipe_slug,
  r.private AS recipe_private,
  r.initial_publish_date AS recipe_initial_publish_date,
  r.forked_from AS recipe_forked_from,
  r.featured_revision AS recipe_featured_revision,
  u.id AS user_id,
  u.display_name AS user_display_name,
  u.email AS user_email,
  u.join_date AS user_join_date,
  u.updated_at AS user_updated_at
FROM
  recipes r
JOIN
  users u ON r.author_id = u.id
WHERE
  r.id = $1
LIMIT 1;

-- name: GetRecipeWithForkedFromById :one
SELECT
  r.id AS recipe_id,
  r.slug AS recipe_slug,
  r.private AS recipe_private,
  r.initial_publish_date AS recipe_initial_publish_date,
  r.forked_from AS recipe_forked_from,
  r.featured_revision AS recipe_featured_revision,
  fr.id AS forked_revision_id,
  fr.recipe_id AS forked_recipe_id,
  fr.parent_id AS forked_parent_id,
  fr.recipe_description AS forked_recipe_description,
  fr.change_comment AS forked_change_comment,
  fr.title AS forked_title,
  fr.publish_date AS forked_publish_date
FROM
  recipes r
LEFT JOIN recipe_revisions fr ON r.forked_from = fr.id
WHERE
  r.id = $1
LIMIT 1;

-- name: GetRecipeWithFeaturedRevisionById :one
SELECT
  r.id AS recipe_id,
  r.slug AS recipe_slug,
  r.private AS recipe_private,
  r.initial_publish_date AS recipe_initial_publish_date,
  r.featured_revision AS recipe_featured_revision,
  fr.id AS featured_revision_id,
  fr.recipe_id AS featured_recipe_id,
  fr.parent_id AS featured_parent_id,
  fr.recipe_description AS featured_recipe_description,
  fr.change_comment AS featured_change_comment,
  fr.title AS featured_title,
  fr.publish_date AS featured_publish_date
FROM
  recipes r
LEFT JOIN recipe_revisions fr ON r.featured_revision = fr.id
WHERE
  r.id = $1
LIMIT 1;



