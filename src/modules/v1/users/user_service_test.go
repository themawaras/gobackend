package users

import (
	"testing"

	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/modules/v1/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var dataUser models.Users

	repo.Mock.On("FindAll").Return(&dataUser, nil)
	data, err := service.FindAll()

	users := data.Message

	assert.Equal(t, "OK", users, "Wrong")
	assert.Nil(t, err)
}

func TestFindByUsername(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var modelMock = models.User{
		Username: "toto",
	}

	repo.Mock.On("FindByUsername", "toto").Return(&modelMock, nil)
	data, err := service.FindByUsername("toto")

	result := data.Data.(*models.User)

	assert.Equal(t, "toto", result.Username, "Expect len username = toto")
	assert.Nil(t, err)
}

func TestAdd(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var AddMock = models.User{
		Username: "mario",
		Password: "abcde",
		Name:     "Mario",
		City:     "Jakarta",
	}

	repo.Mock.On("Add", &AddMock).Return(&AddMock, nil)
	data, err := service.Add(&AddMock)

	result := data.Data.(*models.User)

	assert.Equal(t, "Fatimah", result.Name, "Expect len email = lutfi123@gmail.com")
	assert.Nil(t, err)
}
