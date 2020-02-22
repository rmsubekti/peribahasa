package main

import (
	"fmt"
	"net/http"
	"os"
	"peribahasa/app"
	"peribahasa/web"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	app.UseRoute(api)

	website := router.PathPrefix("/").Subrouter()
	web.UseRoute(website)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
