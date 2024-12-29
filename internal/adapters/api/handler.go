package api

import (
	"encoding/json"              // Import the json package to handle JSON encoding and decoding
	"explorer/internal/usecases" // Import the usecases package to interact with the application's logic layer
	"log"                        // Import the log package for logging
	"net/http"                   // Import the net/http package to handle HTTP requests and responses
)

// FetchShapes is an HTTP handler function that returns the shape data for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve shape data.
func FetchShapes(fetchData usecases.FetchData) http.HandlerFunc {
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

// FetchStops is an HTTP handler function that returns the list of stops for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve stops data.
func FetchStops(fetchData usecases.FetchData) http.HandlerFunc {
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

// UpdateLiveData is an HTTP handler function that returns the live data of vehicles for a given route.
// It extracts the route ID from the request query parameters and calls the FetchData service to retrieve live data (vehicles).
func UpdateLiveData(fetchData usecases.FetchData) http.HandlerFunc {
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
