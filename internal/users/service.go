package users

import (
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
)

type Service struct {
	repository Repository
}

func NewService(config util.Config, repository Repository) *Service {
	return &Service{repository}
}
