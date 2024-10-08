// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPattern = `-- name: CreatePattern :one
INSERT INTO patterns (
  user_id, instructions
) VALUES (
  $1, $2
)
RETURNING user_id, id, instructions
`

type CreatePatternParams struct {
	UserID       uuid.UUID
	Instructions *string
}

func (q *Queries) CreatePattern(ctx context.Context, arg CreatePatternParams) (Pattern, error) {
	row := q.db.QueryRow(ctx, createPattern, arg.UserID, arg.Instructions)
	var i Pattern
	err := row.Scan(&i.UserID, &i.ID, &i.Instructions)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, email, hashed_password, salt
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, username, email, hashed_password, salt
`

type CreateUserParams struct {
	Username       string
	Email          string
	HashedPassword string
	Salt           string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.Salt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Salt,
	)
	return i, err
}

const getPatternById = `-- name: GetPatternById :one
SELECT user_id, id, instructions FROM patterns
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPatternById(ctx context.Context, id uuid.UUID) (Pattern, error) {
	row := q.db.QueryRow(ctx, getPatternById, id)
	var i Pattern
	err := row.Scan(&i.UserID, &i.ID, &i.Instructions)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, hashed_password, salt FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Salt,
	)
	return i, err
}
