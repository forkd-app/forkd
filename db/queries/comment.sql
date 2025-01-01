-- name: ListCommentsByAuthor :many
SELECT
  id,
  recipe_id,
  author_id,
  content,
  post_date
FROM
  recipe_comments
WHERE
  author_id = $1 AND post_date > $2
ORDER BY post_date
LIMIT $3;
