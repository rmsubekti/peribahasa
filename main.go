package main

import (
	"fmt"
	"net/http"
	"os"
	"peribahasa/app/controllers"
	"peribahasa/web/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/register", controllers.Register).Methods("POST")
	api.HandleFunc("/login", controllers.Login).Methods("POST")
	api.HandleFunc("/", controllers.GetRandomPeribahasa).Methods("GET")

	api.HandleFunc("/jenis", controllers.CreateJenis).Methods("POST")
	api.HandleFunc("/jenis/{id}", controllers.GetJenis).Methods("GET")
	api.HandleFunc("/jenis/{id}", controllers.UpdateJenis).Methods("PUT")
	api.HandleFunc("/jenis/{id}", controllers.DeleteJenis).Methods("DELETE")

	api.HandleFunc("/asal", controllers.CreateAsal).Methods("POST")
	api.HandleFunc("/asal/{id}", controllers.GetAsal).Methods("GET")
	api.HandleFunc("/asal/{id}", controllers.UpdateAsal).Methods("PUT")
	api.HandleFunc("/asal/{id}", controllers.DeleteAsal).Methods("DELETE")

	api.HandleFunc("/peribahasa", controllers.CreatePeribahasa).Methods("POST")
	api.HandleFunc("/peribahasa/{id}", controllers.GetPeribahasa).Methods("GET")
	api.HandleFunc("/peribahasa/{id}", controllers.UpdatePeribahasa).Methods("PUT")
	api.HandleFunc("/peribahasa/{id}", controllers.DeletePeribahasa).Methods("DELETE")

	web := router.PathPrefix("/").Subrouter()
	staticFileDirectory := http.Dir("./web/assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	web.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	webadmin := router.PathPrefix("/admin").Subrouter()
	webadmin.HandleFunc("/jenis", handlers.JenisListPage).Methods("GET", "POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
