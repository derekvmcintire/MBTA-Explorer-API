package handlers

import (
	"encoding/json"
	"explorer/internal/core"
	"log"
	"net/http"
)

// UpdateLiveData is an HTTP handler function that returns the live data of vehicles for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve live data (vehicles).
func UpdateLivePosition(fetchData core.MbtaApiHelper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the route ID from the query parameters of the URL (e.g., /live-data?route_id=Red)
		routeID := r.URL.Query().Get("route_id")

		// Call the GetLiveData method of the fetchData service to get the live data for the given route ID
		vehicles, err := fetchData.GetLiveData(routeID)

		// If an error occurred while fetching the live data, log the error and return a 500 Internal Server Error
		if err != nil {
			log.Println("Error in UpdateLiveData:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response header to specify that the content being returned is in JSON format
		w.Header().Set("Content-Type", "application/json")

		// Encode the vehicles data as JSON and send it in the response body
		json.NewEncoder(w).Encode(vehicles)
	}
}
