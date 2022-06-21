package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func subRoute() {
	mainRoute := mux.NewRouter()

	users := mainRoute.PathPrefix("/user").Subrouter()
	users.HandleFunc("/", userHandler)

	products := mainRoute.PathPrefix("/product").Subrouter()
	products.HandleFunc("/", productHandler)

	vehicles := mainRoute.PathPrefix("/vehicle").Subrouter()
	vehicles.HandleFunc("/", vehiclesHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from USER"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from PRODUCT"))
}

func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from VEHICLE"))
}
