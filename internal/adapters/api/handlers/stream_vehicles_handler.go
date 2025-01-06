package handlers

import (
	"explorer/internal/constants"
	"explorer/internal/infrastructure/config"
	"explorer/internal/infrastructure/stream"
	"explorer/internal/usecases"
	"log"
	"net/http"
)

func StreamVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request to stream vehicles")

	// Initialize dependencies
	sm := stream.NewStreamManager()
	useCase := usecases.NewStreamVehiclesUseCase(sm)

	// Set up SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Set up the stream and get client channel
	clientChan := useCase.StreamSetup(
		constants.MbtaVehicleLiveStreamUrl,
		config.GetAPIKey(),
	)
	defer sm.RemoveClient(clientChan)

	// Handle client disconnection
	useCase.HandleDisconnect(r.Context(), clientChan)

	// Stream data to client
	flusher := w.(http.Flusher)
	for data := range clientChan {
		_, _ = w.Write([]byte(data))
		flusher.Flush()
	}
}
