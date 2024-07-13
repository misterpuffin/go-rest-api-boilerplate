package users

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	db "github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
)

type ServiceTestSuite struct {
	suite.Suite
	service          Service
	mockedRepository *MockedRepository
}

type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockedRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	args := m.Called(ctx, params)
	user := db.User{ID: uuid.New(), Username: params.Username, Email: params.Email, HashedPassword: params.HashedPassword, Salt: params.Salt}
	return user, args.Error(1)
}

func (s *ServiceTestSuite) SetupSuite() {
	mockedRepository := MockedRepository{}
	s.service = *NewService(*config.LoadTestConfig(), &mockedRepository)
	s.mockedRepository = &mockedRepository
}

func (suite *ServiceTestSuite) TestRegisterUser() {
	username := "mistermuffin"
	email := "snghaoren@gmail.com"
	password := "password"

	suite.mockedRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(db.User{}, pgx.ErrNoRows)
	suite.mockedRepository.On("CreateUser", mock.Anything, mock.Anything).Return(nil, nil)

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
