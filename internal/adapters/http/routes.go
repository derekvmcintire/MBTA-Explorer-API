package http

import (
	"explorer/internal/adapters/api/handlers"
	"explorer/internal/core"
	"explorer/internal/infrastructure/stream"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, mbtaApiHelper core.MbtaApiHelper, sm *stream.StreamManager) {

	streamHandler := handlers.NewStreamVehiclesHandler(sm)
	sotpsHandler := handlers.FetchRouteStops(mbtaApiHelper)
	shapesHandler := handlers.FetchRouteShapes(mbtaApiHelper)
	routesHandler := handlers.FetchRoutes(mbtaApiHelper)
	livePositionHandler := handlers.UpdateLivePosition(mbtaApiHelper)

	router.Handle("/stream/vehicles", streamHandler)
	router.HandleFunc("/api/stops", sotpsHandler).Methods("GET")
	router.HandleFunc("/api/shapes", shapesHandler).Methods("GET")
	router.HandleFunc("/api/routes", routesHandler).Methods("GET")
	router.HandleFunc("/api/live", livePositionHandler).Methods("GET")
}
