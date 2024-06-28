package users

import (
	"context"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
}
