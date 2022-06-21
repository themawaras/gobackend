package main

import (
	"log"
	"os"

	"github.com/backendGo-learn/src/configs/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

// VERSI LAWAS TANPA SUB ROUTER
// func main() {
// 	mainRoute := mux.NewRouter()

// 	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
// 	mainRoute.HandleFunc("/user/{id}", paramsHandler).Methods("GET")
// 	mainRoute.HandleFunc("/product", queryHandler).Methods("GET")
// 	mainRoute.HandleFunc("/vehicle", bodyHandler).Methods("POST")

// 	if err := http.ListenAndServe(":8088", mainRoute); err != nil {
// 		log.Fatal("Error happend")
// 	}
// }

// func sampleHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello from HOME"))
// }

// func paramsHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	fmt.Fprintf(w, "params : %v", vars["id"])
// }

// func queryHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := r.URL.Query()

// 	fmt.Fprintf(w, "params : %v", vars["name"])
// }

// func bodyHandler(w http.ResponseWriter, r *http.Request) {
// 	var vehicle vehicles
// 	json.NewDecoder(r.Body).Decode(&vehicle)

// 	fmt.Println(vehicle.Name)

// 	// fmt.Fprintf(w, "params : %v", vars["name"])
// }
