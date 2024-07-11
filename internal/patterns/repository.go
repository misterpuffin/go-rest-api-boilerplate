package patterns

import (
	"context"

	"github.com/google/uuid"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type Repository interface {
	GetPatternById(context.Context, uuid.UUID) (db.Pattern, error)
	CreatePattern(context.Context, db.CreatePatternParams) (db.Pattern, error)
}
