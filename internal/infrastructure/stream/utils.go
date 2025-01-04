package stream

import (
	"context"
	"net/http"
)

// createRequest creates an HTTP request for streaming data from the MBTA API.
// It sets the appropriate headers for the request, including the API key and content type.
func (sm *StreamManager) createRequest(ctx context.Context, url, apiKey string) (*http.Request, error) {
	// Create a new HTTP GET request with the provided context, URL, and no body (nil)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		// Return nil and the error if request creation fails
		return nil, err
	}
	// Set the necessary headers for the streaming request
	req.Header.Set("Accept", "text/event-stream") // Set the content type for streaming data
	req.Header.Set("X-API-Key", apiKey)           // Set the API key in the request header
	// Return the created request and nil error
	return req, nil
}

// AddClient adds a new client channel to the manager to receive data updates.
// It locks the client list to ensure thread safety during modification.
func (sm *StreamManager) AddClient(client chan string) {
	sm.clientsMutex.Lock()          // Lock to ensure safe access to the clients map
	defer sm.clientsMutex.Unlock()  // Unlock once the operation is done
	sm.clients[client] = struct{}{} // Add the client channel to the map of active clients
}

// RemoveClient removes a client channel when they disconnect.
// It locks the client list to ensure thread safety during modification.
func (sm *StreamManager) RemoveClient(client chan string) {
	sm.clientsMutex.Lock()         // Lock to ensure safe access to the clients map
	defer sm.clientsMutex.Unlock() // Unlock once the operation is done
	// Check if the client exists in the map
	if _, ok := sm.clients[client]; ok {
		// Remove the client from the map and close the channel to signal disconnection
		delete(sm.clients, client)
		close(client)
	}
}

// Stop stops the stream manager and signals all processes to stop.
// It closes the stop channel to initiate the shutdown process.
func (sm *StreamManager) Stop() {
	// Close the stop channel to signal all goroutines to stop
	close(sm.stop)
}
