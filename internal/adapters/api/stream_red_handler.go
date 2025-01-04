package api

import (
	"explorer/internal/stream"
	"net/http"
)

func StreamVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientChan := make(chan string, 100)
	stream.MBTAStreamManager.AddClient(clientChan)

	defer stream.MBTAStreamManager.RemoveClient(clientChan)

	// Handle client disconnect
	ctx := r.Context()
	go func() {
		<-ctx.Done()
		stream.MBTAStreamManager.RemoveClient(clientChan)
	}()

	flusher := w.(http.Flusher)

	// Stream data to the client
	for data := range clientChan {
		_, _ = w.Write([]byte(data))
		flusher.Flush()
	}
}
