package app

import (
	"peribahasa/app/controllers"

	"github.com/gorilla/mux"
)

// UseRoute api
func UseRoute(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/", controllers.GetRandomPeribahasa).Methods("GET")

	r.HandleFunc("/jenis", controllers.CreateJenis).Methods("POST")
	r.HandleFunc("/jenis/{id}", controllers.GetJenis).Methods("GET")
	r.HandleFunc("/jenis/{id}", controllers.UpdateJenis).Methods("PUT")
	r.HandleFunc("/jenis/{id}", controllers.DeleteJenis).Methods("DELETE")

	r.HandleFunc("/asal", controllers.CreateAsal).Methods("POST")
	r.HandleFunc("/asal/{id}", controllers.GetAsal).Methods("GET")
	r.HandleFunc("/asal/{id}", controllers.UpdateAsal).Methods("PUT")
	r.HandleFunc("/asal/{id}", controllers.DeleteAsal).Methods("DELETE")

	r.HandleFunc("/peribahasa", controllers.CreatePeribahasa).Methods("POST")
	r.HandleFunc("/peribahasa/{id}", controllers.GetPeribahasa).Methods("GET")
	r.HandleFunc("/peribahasa/{id}", controllers.UpdatePeribahasa).Methods("PUT")
	r.HandleFunc("/peribahasa/{id}", controllers.DeletePeribahasa).Methods("DELETE")
}
