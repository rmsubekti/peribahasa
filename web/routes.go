package web

import (
	"net/http"
	"peribahasa/web/handlers"

	"github.com/gorilla/mux"
)

//UseRoute website
func UseRoute(r *mux.Router) {
	staticFileDirectory := http.Dir("./web/assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	webadmin := r.PathPrefix("/admin").Subrouter()
	webadmin.HandleFunc("/jenis", handlers.JenisIndexPage).Methods("GET")
	webadmin.HandleFunc("/jenis/baru", handlers.AddNewJenisPage).Methods("GET", "POST")
	webadmin.HandleFunc("/jenis/ubah/{id}", handlers.EditJenisPage).Methods("GET", "POST")
	webadmin.HandleFunc("/jenis/hapus/{id}", handlers.DeleteJenisPage).Methods("GET", "POST")

	webadmin.HandleFunc("/asal", handlers.AsalIndexPage).Methods("GET")
	webadmin.HandleFunc("/asal/baru", handlers.AddNewAsalPage).Methods("GET", "POST")
	webadmin.HandleFunc("/asal/ubah/{id}", handlers.EditAsalPage).Methods("GET", "POST")
	webadmin.HandleFunc("/asal/hapus/{id}", handlers.DeleteAsalPage).Methods("GET", "POST")

	webadmin.HandleFunc("/peribahasa", handlers.PeribahasaIndexPage).Methods("GET")
	webadmin.HandleFunc("/peribahasa/baru", handlers.AddNewPeribahasaPage).Methods("GET", "POST")
	webadmin.HandleFunc("/peribahasa/ubah/{id}", handlers.EditPeribahasaPage).Methods("GET", "POST")
	webadmin.HandleFunc("/peribahasa/hapus/{id}", handlers.DeletePeribahasaPage).Methods("GET", "POST")
}
