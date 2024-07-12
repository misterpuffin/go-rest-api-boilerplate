package users

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	db "github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service Service
}

type MockRepository struct{}

func (r MockRepository) GetUserByEmail(_ context.Context, _ string) (db.User, error) {
	return db.User{}, pgx.ErrNoRows
}
func (r MockRepository) CreateUser(_ context.Context, params db.CreateUserParams) (db.User, error) {
	return db.User{ID: uuid.New(), Username: params.Username, Email: params.Email, HashedPassword: params.HashedPassword, Salt: params.Salt}, nil
}

func (s *ServiceTestSuite) SetupSuite() {
	mockRepository := MockRepository{}
	s.service = *NewService(*config.LoadTestConfig(), mockRepository)
}

func (suite *ServiceTestSuite) TestRegisterUser() {
	username := "mistermuffin"
	email := "snghaoren@gmail.com"
	password := "password"
	user, err := suite.service.RegisterUser(username, email, password)
	suite.Assert().NoError(err)
	suite.Assert().NotEmpty(user.ID)
	suite.Assert().Contains(user.Username, username)
}

func (suite *ServiceTestSuite) TestLoginUser() {
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
