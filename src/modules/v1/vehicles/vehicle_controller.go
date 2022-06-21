package vehicles

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	repo *vehicle_repo
}

func NewCtrl(rep *vehicle_repo) *vehicle_ctrl {
	return &vehicle_ctrl{rep}
}

func (rep *vehicle_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.Add(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *vehicle_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data Vehicle

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		log.Fatalf("Cannot convert string to integer. %v", err)
	}

	result, err := rep.repo.Update(&id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *vehicle_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])
	if err != nil {
		log.Fatalf("Cannot convert string to integer.  %v", err)
	}

	result, err := rep.repo.Delete(&id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *vehicle_ctrl) SortByYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataYear = r.URL.Query()
	year, err := strconv.Atoi(dataYear["year"][0])
	if err != nil {
		log.Fatalf("Cannot convert string to integer.  %v", err)
	}

	data, err := rep.repo.SortByYear(&year)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}
