package api

import (
	"encoding/json"
	"explorer/internal/ports/core"
	"log"
	"net/http"
)

// FetchRouteStops is an HTTP handler function that returns the list of stops for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve stops data.
func FetchRouteStops(fetchData core.FetchFromMBTAUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the route ID from the query parameters of the URL (e.g., /stops?route_id=Red)
		routeID := r.URL.Query().Get("route_id")

		// Call the GetStops method of the fetchData service to get the stops for the given route ID
		stops, err := fetchData.GetStops(routeID)

		// If an error occurred while fetching the stops, log the error and return a 500 Internal Server Error
		if err != nil {
			log.Println("Error in FetchStops:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response header to specify that the content being returned is in JSON format
		w.Header().Set("Content-Type", "application/json")

		// Encode the stops data as JSON and send it in the response body
		json.NewEncoder(w).Encode(stops)
	}
}
