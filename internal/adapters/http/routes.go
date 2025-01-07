package http

import (
	"explorer/internal/adapters/api/handlers"
	"explorer/internal/core"
	"explorer/internal/infrastructure/stream"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, mbtaApiHelper core.MbtaApiHelper, sm *stream.StreamManager) {

	streamHandler := handlers.NewStreamVehiclesHandler(sm)

	router.HandleFunc("/api/stops", handlers.FetchRouteStops(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/shapes", handlers.FetchRouteShapes(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/live", handlers.UpdateLivePosition(mbtaApiHelper)).Methods("GET")
	router.HandleFunc("/api/routes", handlers.FetchRoutes(mbtaApiHelper)).Methods("GET")
	router.Handle("/stream/vehicles", streamHandler)
}
