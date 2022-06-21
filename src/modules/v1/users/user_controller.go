package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/backendGo-learn/src/database/gorm/models"
	"github.com/backendGo-learn/src/interfaces"
	"github.com/gorilla/mux"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(ctrl interfaces.UserService) *user_ctrl {
	return &user_ctrl{ctrl}
}

func (ctrl *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.FindAll()
	if err != nil {
		response.ResponseJSON(400, "failed show data").Send(w)
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *user_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	result, err := ctrl.svc.Add(&data)
	if err != nil {
		response.ResponseJSON(400, "failed add data").Send(w)
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])

	if err != nil {
		response.ResponseJSON(400, "string to in convert error").Send(w)
	}

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		response.ResponseJSON(400, "failed delete data").Send(w)
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.User

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		response.ResponseJSON(400, "string to in convert error").Send(w)
		return
	}

	result, err := ctrl.svc.Update(id, &data)
	if err != nil {
		response.ResponseJSON(400, "update data failed").Send(w)
	}

	json.NewEncoder(w).Encode(&result)
}
