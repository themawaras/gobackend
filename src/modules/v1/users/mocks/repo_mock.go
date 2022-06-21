package mocks

import (
	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (rm *RepoMock) FindAll() (*models.Users, error) {
	args := rm.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (rm *RepoMock) FindByUsername(username string) (*models.User, error) {
	args := rm.Mock.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (rm *RepoMock) Add(data *models.User) (*models.User, error) {
	args := rm.Mock.Called(data)
	return args.Get(0).(*models.User), args.Error(1)
}

func (rm *RepoMock) Update(id int, data *models.User) (*models.User, error) {
	args := rm.Mock.Called(id, data)
	return args.Get(0).(*models.User), args.Error(1)
}
