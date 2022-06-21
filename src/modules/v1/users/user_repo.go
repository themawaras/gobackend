package users

import (
	"encoding/json"
	"errors"

	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/helpers"
	"gorm.io/gorm"
)

var response helpers.Response

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (repo *user_repo) FindAll() (*models.Users, error) {

	var users models.Users

	result := repo.db.Order("user_id desc").Find(&users)

	if result.Error != nil {
		return nil, errors.New("failed show the data")
	}

	return &users, nil
}

func (repo *user_repo) Add(data *models.User) (*models.User, error) {

	var users models.User

	getUsername := repo.db.Where("username = ?", &data.Username).First(&users)
	if getUsername.RowsAffected != 0 {
		return nil, errors.New("Username has been registered")
	}

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Failed to add data")
	}

	getdata := repo.db.First(&users, &data.User_id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("Username has been registered")
	}

	return &users, nil
}

func (repo *user_repo) Update(id int, data *models.User) (*models.User, error) {

	var users *models.User

	result := repo.db.Model(&models.User{}).Where("user_id = ?", &id).Updates(&models.User{Username: data.Username, Name: data.Name, City: data.City})

	if result.Error != nil {
		return nil, errors.New("Failed to update the data")
	}

	getdata := repo.db.First(&users, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("Data not found")
	}

	return users, nil
}

func (repo *user_repo) Delete(id int) (*models.User, error) {

	var users models.User

	getdata := repo.db.First(&users, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data not found")
	}

	result := repo.db.Delete(&models.User{}, &id)

	if result.Error != nil {
		return nil, errors.New("data delete failed")
	}

	return &users, nil
}

func (repo *user_repo) FindByUsername(username string) (*models.User, error) {

	var users models.User

	result := repo.db.First(&users, "username = ?", username)

	if result.RowsAffected < 1 {
		err := json.Unmarshal([]byte("username"), &users)
		return nil, err
	}

	if result.Error != nil {
		err := json.Unmarshal([]byte("Failed to retrieve data"), &users)
		return nil, err
	}

	return &users, nil
}

// func (repo *user_repo) SortByCity(city string) (*helpers.Response, error) {

// 	var users models.Users

// 	result := repo.db.Order("city desc").Where("CAST(city AS string) > ?", city).Find(&users)

// 	if result.RowsAffected < 1 {
// 		res := response.ResponseJSON(404, users)
// 		return res, nil
// 	}

// 	if result.Error != nil {
// 		res := response.ResponseJSON(500, users)
// 		return res, nil
// 	}

// 	res := response.ResponseJSON(200, users)
// 	return res, nil
// }
