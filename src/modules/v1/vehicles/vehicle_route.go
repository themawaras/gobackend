package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/add", ctrl.AddData).Methods("POST")
	route.HandleFunc("/update", ctrl.UpdateData).Methods("PUT")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/year", ctrl.SortByYear).Methods("GET")
}
