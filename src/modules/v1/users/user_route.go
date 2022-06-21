package users

import (
	"github.com/backendGo-learn/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/add", middleware.Do(ctrl.AddData, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Update, middleware.CheckAuth)).Methods("DELETE")
	// route.HandleFunc("/find", ctrl.FindByUsername).Methods("POST")
}
