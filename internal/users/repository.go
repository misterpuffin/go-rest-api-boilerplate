package users

import (
	"context"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type Repository interface {
	GetUserByEmail(context.Context, string) (db.User, error)
	CreateUser(context.Context, db.CreateUserParams) (db.User, error)
}
