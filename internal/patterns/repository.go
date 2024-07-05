package patterns

import (
	"context"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type Repository interface {
	GetPatternById(context.Context, string) (db.Pattern, error)
	CreatePattern(context.Context, db.CreatePatternParams) (db.Pattern, error)
}
