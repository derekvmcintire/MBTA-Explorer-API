package http

import (
	"explorer/internal/adapters/mbta/api/handlers"
	"explorer/internal/core"
	ports "explorer/internal/ports/streaming"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all the HTTP routes for the application.
//
// Parameters:
// - router: The Gorilla Mux router used to define the HTTP endpoints.
// - mbtaApiHelper: Helper interface for interacting with the MBTA API.
// - sm: StreamManagerUseCase responsible for managing vehicle streaming.
func RegisterRoutes(router *mux.Router, mbtaApiHelper core.MbtaApiHelper, sm ports.StreamManager) {

	// Initialize handlers for each route
	streamHandler := handlers.NewStreamVehiclesHandler(sm)            // Handles streaming of vehicle data
	sotpsHandler := handlers.FetchRouteStops(mbtaApiHelper)           // Handles fetching route stops
	shapesHandler := handlers.FetchRouteShapes(mbtaApiHelper)         // Handles fetching route shapes
	routesHandler := handlers.FetchRoutes(mbtaApiHelper)              // Handles fetching available routes
	livePositionHandler := handlers.UpdateLivePosition(mbtaApiHelper) // Handles live vehicle positions

	// Define HTTP endpoints and their corresponding handlers
	router.Handle("/stream/vehicles", streamHandler)                   // Streaming endpoint for vehicle data
	router.HandleFunc("/api/stops", sotpsHandler).Methods("GET")       // Fetch route stops via GET
	router.HandleFunc("/api/shapes", shapesHandler).Methods("GET")     // Fetch route shapes via GET
	router.HandleFunc("/api/routes", routesHandler).Methods("GET")     // Fetch route list via GET
	router.HandleFunc("/api/live", livePositionHandler).Methods("GET") // Fetch live vehicle positions via GET
}
