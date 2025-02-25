// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: get_recipe_by_slug.sql

package db

import (
	"context"
)

const getRecipeBySlug = `-- name: GetRecipeBySlug :one
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
LIMIT 1
`

func (q *Queries) GetRecipeBySlug(ctx context.Context, slug string) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipeBySlug, slug)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Private,
		&i.InitialPublishDate,
		&i.ForkedFrom,
		&i.FeaturedRevision,
	)
	return i, err
}
