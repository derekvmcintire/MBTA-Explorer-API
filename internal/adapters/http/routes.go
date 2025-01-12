package http

import (
	"explorer/internal/adapters/mbta/api/handlers"
	"explorer/internal/core/usecases"
	"explorer/internal/infrastructure/middleware"
	ports "explorer/internal/ports/streaming"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all the HTTP routes for the application.
//
// Parameters:
// - router: The Gorilla Mux router used to define the HTTP endpoints.
// - mbtaApiHelper: Helper interface for interacting with the MBTA API.
// - sm: StreamManagerUseCase responsible for managing vehicle streaming.
func RegisterRoutes(router *mux.Router, mbtaApiHelper usecases.MbtaApiHelper, sm ports.StreamManager) {

	// Initialize handlers for each route
	streamVehiclesHandler := handlers.NewStreamVehiclesHandler(sm)                                       // Handles streaming of vehicle data
	vehiclePositionHandler := middleware.CompressHandler(handlers.VehiclePositionHandler(mbtaApiHelper)) // Handles live vehicle positions
	routesHandler := middleware.CompressHandler(handlers.RouteHandler(mbtaApiHelper))

	// Define HTTP endpoints and their corresponding handlers
	router.Handle("/stream/vehicles", streamVehiclesHandler)              // Streaming endpoint for vehicle data
	router.Handle("/api/routes", routesHandler).Methods("GET")            // Fetch route list via GET
	router.Handle("/api/vehicles", vehiclePositionHandler).Methods("GET") // Fetch live vehicle positions via GET
}
