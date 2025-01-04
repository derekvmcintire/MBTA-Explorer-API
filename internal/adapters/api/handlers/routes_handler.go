package handlers

import (
	"encoding/json"                           // Import the json package for JSON encoding/decoding
	"explorer/internal/adapters/api/response" // Import response models for structured API responses
	"explorer/internal/core"                  // Import core ports for accessing use case implementations
	"log"                                     // Import the log package for logging errors and information
	"net/http"                                // Import net/http for building HTTP handlers
	"strings"                                 // Import strings package to handle string operations like splitting
)

// FetchRoutes is an HTTP handler function that returns all relevant data (stops and shapes)
// for a list of route IDs provided in the request query parameters.
func FetchRoutes(useCases core.MbtaApiHelper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received...")

		// Extract the route IDs from the query parameter (e.g., /routes?route_ids=Red,Orange,Blue)
		strRouteIDs := r.URL.Query().Get("route_ids")

		log.Println("fetching routes...")
		log.Println("routes are: ", strRouteIDs)

		// Split the comma-separated route IDs into a slice of strings
		routeIDs := strings.Split(strRouteIDs, ",")

		// Initialize a slice to hold the aggregated responses for each route
		var responses []response.GetRouteResponse

		// Iterate through each route ID to fetch stops and shapes
		for _, routeID := range routeIDs {
			// Fetch stops for the current route ID
			stops, err := useCases.GetStops(routeID)
			if err != nil {
				log.Printf("Error fetching stops for route %s: %v", routeID, err)
				http.Error(w, "Error fetching stops for one or more routes", http.StatusInternalServerError)
				return
			}

			// Fetch shapes for the current route ID
			shapes, err := useCases.GetShapes(routeID)
			if err != nil {
				log.Printf("Error fetching shapes for route %s: %v", routeID, err)
				http.Error(w, "Error fetching shapes for one or more routes", http.StatusInternalServerError)
				return
			}

			// Construct the response for the current route
			routeResponse := response.GetRouteResponse{
				ID:          routeID,            // Set the route ID
				Coordinates: shapes.Coordinates, // Assign the decoded shapes (coordinates)
				Stops:       stops,              // Assign the fetched stops
			}

			// Add the constructed response to the list
			responses = append(responses, routeResponse)
		}

		// Set the Content-Type header to indicate JSON response
		w.Header().Set("Content-Type", "application/json")

		// Encode the aggregated responses as JSON and send them in the response body
		if err := json.NewEncoder(w).Encode(responses); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
