package routes

import (
	"explorer/internal/adapters/api"
	"explorer/internal/ports/core"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, mbtaApiHelper core.MbtaApiHelper) {
	router.HandleFunc("/api/stops", api.FetchRouteStops(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/shapes", api.FetchRouteShapes(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/live", api.UpdateLivePosition(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/routes", api.FetchRoutes(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/stream/vehicles", api.StreamVehiclesHandler)
}
