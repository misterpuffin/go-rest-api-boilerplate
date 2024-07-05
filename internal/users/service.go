package users

import (
	"context"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
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

func (s Service) RegisterUser(username string, email string, password string) (user db.User, err error) {
	ctx := context.Background()
	user, err = s.repository.GetUserByEmail(ctx, email)

	if err != pgx.ErrNoRows {
		if err != nil {
			return
		}
		return user, util.BadRequest("Email already registered")
	}

	const saltSize = 16
	salt := util.GenerateRandomSalt(saltSize)
	hashedPassword := util.HashPassword(password, salt)

	params := db.CreateUserParams{Username: username, Email: email, HashedPassword: hashedPassword, Salt: salt}
	return s.repository.CreateUser(ctx, params)
}

func (s Service) LoginUser(email string, password string) (token string, err error) {
	ctx := context.Background()
	user, err := s.repository.GetUserByEmail(ctx, email)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", util.BadRequest("Email has not been registered")
		}
		return
	}

	salt := user.Salt
	checkHashedPassword := util.HashPassword(password, salt)

	if checkHashedPassword != user.HashedPassword {
		return "", util.BadRequest("Incorrect Password")
	}

	token, err = util.CreateJWTToken(util.JWTTokenPayload{UserId: user.ID.String()}, s.config)
	return token, nil
}
