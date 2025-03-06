package service

import (
	"testing"

	"github.com/sabrodigan/newWebBoxes/server/dto"
	"github.com/sabrodigan/newWebBoxes/server/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockUserService is a mock implementation of IUserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	args := m.Called(data, sessionContext)
	return args.Get(0), args.Error(1)
}

func (m *MockUserService) FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error) {
	args := m.Called(email, sessionContext)
	return args.Get(0), args.Error(1)
}

func TestAuthService_Login(t *testing.T) {
	mockUserService := new(MockUserService)
	authService := &AuthService{
		userService: mockUserService,
	}

	loginDto := dto.LoginDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := model.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUserService.On("FindOneUserByEmail", loginDto.Email, mock.Anything).Return(user, nil)

	sessionContext := mongo.SessionContext(nil) // Mock session context

	result, err := authService.Login(loginDto, sessionContext)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result["user"])
	assert.Equal(t, "", result["token"])

	mockUserService.AssertExpectations(t)
}

func TestAuthService_Login_InvalidCredentials(t *testing.T) {
	mockUserService := new(MockUserService)
	authService := &AuthService{
		userService: mockUserService,
	}

	loginDto := dto.LoginDto{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	user := model.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUserService.On("FindOneUserByEmail", loginDto.Email, mock.Anything).Return(user, nil)

	sessionContext := mongo.SessionContext(nil) // Mock session context

	result, err := authService.Login(loginDto, sessionContext)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid credentials")

	mockUserService.AssertExpectations(t)
}
