package auth

import (
	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/helpers"
	"github.com/backendGo-learn/src/interfaces"
)

var response helpers.Response

type token_response struct {
	Tokens string `json:"token"`
}

type Auth_service struct {
	rep interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *Auth_service {
	return &Auth_service{svc}
}

func (svc *Auth_service) Login(body models.User) (*helpers.Response, error) {

	user, err := svc.rep.FindByUsername(body.Username)
	if err != nil {
		return response.ResponseJSON(400, "wrong username"), nil
	}

	if !helpers.CheckPassword(user.Password, body.Password) {
		return response.ResponseJSON(400, "wrong password"), nil
	}

	token := helpers.NewToken(body.Username)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}

	res := response.ResponseJSON(200, token_response{Tokens: theToken})
	return res, nil
}
