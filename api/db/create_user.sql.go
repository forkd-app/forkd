// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: create_user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
  users (
    email,
    display_name
  )
VALUES (
  $1,
  $2
)
RETURNING
  users.id,
  users.display_name,
  users.email,
  users.join_date,
  users.updated_at,
  users.photo
`

type CreateUserParams struct {
	Email       string
	DisplayName string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.DisplayName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
		&i.Photo,
	)
	return i, err
}
