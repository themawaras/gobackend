package vehicles

import (
	"github.com/backendGo-learn/src/helpers"

	"gorm.io/gorm"
)

var vehicles Vehicles
var response helpers.Response

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}
}

func (r *vehicle_repo) FindAll() (*helpers.Response, error) {

	result := r.db.Order("vehicle_id desc").Find(&vehicles)

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) Add(data *Vehicle) (*helpers.Response, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	getdata := r.db.First(&vehicles, &data.Vehicle_id)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(201, vehicles)
	return res, nil
}

func (r *vehicle_repo) Update(id *int, data *Vehicle) (*helpers.Response, error) {

	result := r.db.Model(&Vehicle{}).Where("vehicle_id = ?", &id).Updates(&Vehicle{Manufacture: data.Manufacture, Model: data.Model, Type: data.Type, Color: data.Color, Year: data.Year})

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	getdata := r.db.First(&vehicles, &id)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(201, vehicles)
	return res, nil
}

func (r *vehicle_repo) Delete(data *int) (*helpers.Response, error) {

	getdata := r.db.First(&vehicles, &data)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	result := r.db.Delete(&Vehicle{}, &data)

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) SortByYear(year *int) (*helpers.Response, error) {

	result := r.db.Order("year desc").Where("CAST(year AS int) > ?", year).Find(&vehicles)

	if result.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}
