package patterns

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/errors"
)

type Service struct {
	config     config.Config
	repository Repository
}

func NewService(config config.Config, repository Repository) *Service {
	return &Service{config, repository}
}

func (s Service) GetPattern(patternId string) (pattern db.Pattern, err error) {
	ctx := context.Background()
	parsedId, err := uuid.Parse(patternId)
	if err != nil {
		return db.Pattern{}, errors.BadRequest("Invalid ID")
	}
	pattern, err = s.repository.GetPatternById(ctx, parsedId)

	if err != pgx.ErrNoRows {
		return db.Pattern{}, errors.BadRequest(fmt.Sprintf("Pattern with id %s not found", patternId))
	}
	return pattern, nil
}

func (s Service) CreatePattern(userId uuid.UUID, instructions string) (pattern db.Pattern, err error) {
	ctx := context.Background()
	pattern, err = s.repository.CreatePattern(ctx, db.CreatePatternParams{UserID: userId, Instructions: &instructions})
	return pattern, nil
}
