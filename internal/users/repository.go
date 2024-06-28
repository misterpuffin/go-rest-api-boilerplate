package users

import (
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type Repository interface {
	GetUserByEmail(email string) db.User
}
