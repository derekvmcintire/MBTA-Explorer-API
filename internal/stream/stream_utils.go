package stream

import (
	"context"
	"net/http"
)

// createRequest creates an HTTP request for streaming data from the MBTA API.
func (sm *StreamManager) createRequest(ctx context.Context, url, apiKey string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("X-API-Key", apiKey)
	return req, nil
}

// AddClient adds a new client channel to the manager to receive data updates.
func (sm *StreamManager) AddClient(client chan string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	sm.clients[client] = struct{}{} // Add the client channel to the map
}

// RemoveClient removes a client channel when they disconnect.
func (sm *StreamManager) RemoveClient(client chan string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	if _, ok := sm.clients[client]; ok {
		delete(sm.clients, client) // Remove the client from the map
		close(client)              // Close the channel to signal disconnection
	}
}

// Stop stops the stream manager and signals all processes to stop.
func (sm *StreamManager) Stop() {
	close(sm.stop) // Close the stop channel to signal shutdown
}
