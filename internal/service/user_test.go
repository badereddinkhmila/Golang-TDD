package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userRepoMock struct {
	mock.Mock
}

func (m *userRepoMock) CreateUser() {}

type userServiceSuite struct {
	suite.Suite
}

func (s *userServiceSuite) TestCreateUser()

func TestSuite(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}
