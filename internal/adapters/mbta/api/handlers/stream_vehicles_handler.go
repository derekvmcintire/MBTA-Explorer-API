package handlers

import (
	"explorer/internal/constants"
	"explorer/internal/infrastructure/config"
	"explorer/internal/infrastructure/stream"
	"explorer/internal/usecases"
	"net/http"
)

type StreamVehiclesHandler struct {
	streamManager *stream.StreamManager
	useCase       *usecases.StreamVehiclesUseCase
}

func NewStreamVehiclesHandler(sm *stream.StreamManager) *StreamVehiclesHandler {
	return &StreamVehiclesHandler{
		streamManager: sm,
		useCase:       usecases.NewStreamVehiclesUseCase(sm),
	}
}

// Changed method name from StreamVehiclesHandler to ServeHTTP
func (h *StreamVehiclesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set up SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Set up the stream and get client channel
	clientChan := h.useCase.StreamSetup(
		constants.MbtaVehicleLiveStreamUrl,
		config.GetAPIKey(),
	)
	defer h.streamManager.RemoveClient(clientChan)

	// Handle client disconnection
	h.useCase.HandleDisconnect(r.Context(), clientChan)

	// Stream data to client
	flusher := w.(http.Flusher)
	for data := range clientChan {
		_, _ = w.Write([]byte(data))
		flusher.Flush()
	}
}
