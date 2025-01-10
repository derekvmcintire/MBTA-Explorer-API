package handlers

import (
	"explorer/internal/constants"
	"explorer/internal/core/usecases"
	"explorer/internal/infrastructure/config"
	ports "explorer/internal/ports/streaming"
	"net/http"
)

// StreamVehiclesHandler is responsible for handling the streaming of vehicle data
// via Server-Sent Events (SSE).
type StreamVehiclesHandler struct {
	streamManager ports.StreamManager             // Manages the global streaming state
	useCase       *usecases.StreamVehiclesUseCase // Use case for vehicle streaming logic
}

// NewStreamVehiclesHandler creates a new instance of StreamVehiclesHandler.
//
// Parameters:
// - sm: The StreamManagerUseCase instance to manage vehicle streams.
//
// Returns:
// - A pointer to the initialized StreamVehiclesHandler.
func NewStreamVehiclesHandler(sm ports.StreamManager) *StreamVehiclesHandler {
	return &StreamVehiclesHandler{
		streamManager: sm,
		useCase:       usecases.NewStreamVehiclesUseCase(sm), // Initialize the streaming use case
	}
}

// ServeHTTP implements the http.Handler interface and handles SSE streaming.
//
// Parameters:
// - w: The HTTP response writer to send SSE events to the client.
// - r: The HTTP request object.
//
// Functionality:
// - Sets up necessary SSE headers.
// - Initializes the streaming setup and retrieves a client channel.
// - Listens for and sends data updates to the client until the connection is closed.
func (h *StreamVehiclesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set SSE-specific headers to enable a persistent connection for streaming data.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Initialize the stream and obtain a dedicated channel for this client.
	clientChan := h.useCase.StreamSetup(
		constants.MbtaVehicleLiveStreamUrl, // URL for the MBTA vehicle live stream
		config.GetAPIKey(),                 // API key for authentication
	)
	defer h.streamManager.RemoveClient(clientChan) // Ensure client is removed when function exits.

	// Handle client disconnection to prevent resource leaks.
	h.useCase.HandleDisconnect(r.Context(), clientChan)

	// Stream data to the client as it becomes available.
	flusher := w.(http.Flusher) // Ensure the response writer supports flushing.
	for data := range clientChan {
		_, _ = w.Write([]byte(data)) // Send data to the client
		flusher.Flush()              // Ensure data is immediately sent
	}
}
