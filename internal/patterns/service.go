package patterns

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
)

type Service struct {
	config     util.Config
	repository Repository
}

func NewService(config util.Config, repository Repository) *Service {
	return &Service{config, repository}
}

func (s Service) GetPattern(patternId string) (pattern db.Pattern, err error) {
	ctx := context.Background()
	pattern, err = s.repository.GetPatternById(ctx, patternId)

	if err != pgx.ErrNoRows {
		return db.Pattern{}, util.BadRequest(fmt.Sprintf("Pattern with id %s not found", patternId))
	}
	return pattern, nil
}

func (s Service) CreatePattern(instructions string) (pattern db.Pattern, err error) {
	ctx := context.Background()
	userId, err := uuid.NewRandom()
	pattern, err = s.repository.CreatePattern(ctx, db.CreatePatternParams{UserID: userId, Instructions: &instructions})
	return pattern, nil
}
