package interfaces

import (
	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/helpers"
)

type AuthService interface {
	Login(body models.User) (*helpers.Response, error)
}
