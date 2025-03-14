// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: create_magic_link.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMagicLink = `-- name: CreateMagicLink :one
INSERT INTO
  magic_links (
    user_id,
    token,
    expiry
  )
VALUES (
  $1,
  $2,
  $3
)
RETURNING
  magic_links.id,
  magic_links.token
`

type CreateMagicLinkParams struct {
	UserID pgtype.UUID
	Token  pgtype.UUID
	Expiry pgtype.Timestamp
}

type CreateMagicLinkRow struct {
	ID    pgtype.UUID
	Token pgtype.UUID
}

func (q *Queries) CreateMagicLink(ctx context.Context, arg CreateMagicLinkParams) (CreateMagicLinkRow, error) {
	row := q.db.QueryRow(ctx, createMagicLink, arg.UserID, arg.Token, arg.Expiry)
	var i CreateMagicLinkRow
	err := row.Scan(&i.ID, &i.Token)
	return i, err
}
