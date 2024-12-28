package main

import (
	"log"
	"net/http"

	"github.com/derekvmcintire/MBTA-Explorer-API/internal/adapters/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Initialize use cases
	fetchData := usecases.NewFetchData(usecases.NewMBTAAPIClient())

	// Set up API routes
	r.HandleFunc("/api/routes", api.FetchRoutes(fetchData)).Methods("GET")
	r.HandleFunc("/api/stops", api.FetchStops(fetchData)).Methods("GET")
	r.HandleFunc("/api/live", api.UpdateLiveData(fetchData)).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
