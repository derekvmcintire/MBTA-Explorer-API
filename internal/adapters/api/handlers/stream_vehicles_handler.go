package handlers

import (
	"explorer/internal/constants"
	"explorer/internal/infrastructure/config"
	"explorer/internal/infrastructure/stream"
	"log"
	"net/http"
)

// StreamVehiclesHandler handles the streaming of vehicle data to clients via Server-Sent Events (SSE).
// This function sets up SSE headers, registers the client with the StreamManager,
// and streams data until the client disconnects.
func StreamVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request to stream vehicles")
	sm := stream.NewStreamManager()

	url := constants.MbtaVehicleLiveStreamUrl

	apiKey := config.GetAPIKey()

	sm.EnsureStreaming(url, apiKey)

	// Set headers to establish the response as an SSE stream.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a buffered channel for this client togit receive data.
	clientChan := make(chan string, 100)

	// Register the client channel with the StreamManager.
	sm.AddClient(clientChan)

	// Ensure the client is removed from the StreamManager when the function exits.
	defer sm.RemoveClient(clientChan)

	// Monitor for client disconnection and remove the client when disconnected.
	ctx := r.Context()
	go func() {
		<-ctx.Done() // Wait for the context to signal cancellation.
		sm.RemoveClient(clientChan)
	}()

	// Assert that the ResponseWriter supports the http.Flusher interface for streaming.
	flusher := w.(http.Flusher)

	// Continuously stream data from the client channel to the response.
	for data := range clientChan {
		// Write the data to the HTTP response.
		_, _ = w.Write([]byte(data))

		// Flush the data immediately to ensure the client receives it in real-time.
		flusher.Flush()
	}
}
