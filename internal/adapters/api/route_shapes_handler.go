package api

import (
	"encoding/json"
	"explorer/internal/ports/core"
	"log"
	"net/http"
)

// FetchRouteShapes is an HTTP handler function that returns the shape data for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve shape data.
func FetchRouteShapes(fetchData core.MbtaApiHelper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the route ID from the query parameters of the URL (e.g., /shapes?route_id=Red)
		routeID := r.URL.Query().Get("route_id")

		// Call the GetShapes method of the fetchData service to get the shapes for the given route ID
		shapes, err := fetchData.GetShapes(routeID)

		// If an error occurred while fetching the shapes, log the error and return a 500 Internal Server Error
		if err != nil {
			log.Println("Error in FetchShapes:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response header to specify that the content being returned is in JSON format
		w.Header().Set("Content-Type", "application/json")

		// Encode the shapes data as JSON and send it in the response body
		json.NewEncoder(w).Encode(shapes)
	}
}
