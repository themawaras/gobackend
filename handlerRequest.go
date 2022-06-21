package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type vehicles struct {
	Name string
	Type string
	Year int
}

func handlerReqs() {
	mainRoute := mux.NewRouter()

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc("/user/{id}", paramsHandler).Methods("GET")
	mainRoute.HandleFunc("/product", queryHandler).Methods("GET")
	mainRoute.HandleFunc("/vehicle", bodyHandler).Methods("POST")
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from HOME"))
}

func paramsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "params : %v", vars["id"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	fmt.Fprintf(w, "params : %v", vars["name"])
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	var vehicle vehicles
	json.NewDecoder(r.Body).Decode(&vehicle)

	fmt.Println(vehicle.Name)

	// fmt.Fprintf(w, "params : %v", vars["name"])
}
