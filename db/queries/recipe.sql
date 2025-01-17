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
-- name: GetRecipeByRevisionID :one
SELECT
  recipes.id,
  recipes.author_id,
  recipes.slug,
  recipes.private,
  recipes.initial_publish_date,
  recipes.forked_from,
  recipes.featured_revision
FROM
  recipe_revisions
JOIN
  recipes ON recipe_revisions.recipe_id = recipes.id
WHERE
  recipes.id = $1
LIMIT 1;
-- name: GetRecipeRevisionByParentID :one
SELECT
  parent.id,
  parent.recipe_id,
  parent.parent_id,
  parent.recipe_description,
  parent.change_comment,
  parent.title,
  parent.publish_date
FROM
  recipe_revisions child
JOIN
  recipe_revisions parent ON child.parent_id = parent.id
WHERE
  recipe_revisions.id = $1
LIMIT 1;
-- name: ListIngredientsByRecipeRevisionID :many
SELECT
  recipe_ingredients.id,
  recipe_ingredients.revision_id,
  recipe_ingredients.ingredient,
  recipe_ingredients.quantity,
  recipe_ingredients.unit,
  recipe_ingredients.comment
FROM
  recipe_revisions
JOIN
  recipe_ingredients ON recipe_revisions.id = recipe_ingredients.revision_id 
WHERE
  recipe_revisions.id = $1
ORDER BY recipe_ingredients.id;
-- name: ListStepsByRecipeRevisionID :many
SELECT
  recipe_steps.id,
  recipe_steps.revision_id,
  recipe_steps.content,
  recipe_steps.index
FROM
  recipe_revisions
JOIN
  recipe_steps ON recipe_revisions.id = recipe_steps.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY
  recipe_steps.id;
