package routers

import (
	"net/http"

	"github.com/backendGo-learn/src/configs/database"
	"github.com/backendGo-learn/src/modules/v1/auth"
	"github.com/backendGo-learn/src/modules/v1/users"
	"github.com/backendGo-learn/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	vehicles.New(mainRoute, db)
	users.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from HOME"))
}
