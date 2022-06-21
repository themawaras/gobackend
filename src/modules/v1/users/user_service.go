package users

import (
	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/helpers"
	"github.com/backendGo-learn/src/interfaces"

	"github.com/asaskevich/govalidator"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *user_service {
	return &user_service{svc}
}

func (svc *user_service) FindAll() (*helpers.Response, error) {
	result, err := svc.repo.FindAll()
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *user_service) Add(data *models.User) (*helpers.Response, error) {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	hsPass, err := helpers.HashPassword(data.Password)
	if err != nil {
		res := response.ResponseJSON(400, hsPass)
		return res, nil
	}

	data.Password = hsPass
	result, err := svc.repo.Add(data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *user_service) Update(id int, data *models.User) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Update(id, data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}

func (svc *user_service) Delete(id int) (*helpers.Response, error) {
	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "wrong ID")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Delete(id)
	if err != nil {
		res := response.ResponseJSON(404, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}

func (svc *user_service) FindByUsername(username string) (*helpers.Response, error) {

	result, err := svc.repo.FindByUsername(username)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}

// func (re *user_service) SortByCity(city string) (*help.Response, error) {
// 	data, err := re.rep.SortByCity(city)

// 	if err != nil {
// 		return nil, err
// 	}

// 	response := resp.ResponseJSON(200, data)
// 	return response, nil
// }
