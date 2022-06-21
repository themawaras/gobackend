package interfaces

import (
	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	Add(data *models.User) (*models.User, error)
	Update(id int, data *models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
	// SortByCity(city string) (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
}

type UserService interface {
	FindAll() (*helpers.Response, error)
	Add(data *models.User) (*helpers.Response, error)
	Update(id int, data *models.User) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	// SortByCity(city string) (*helpers.Response, error)
	FindByUsername(username string) (*helpers.Response, error)
}
