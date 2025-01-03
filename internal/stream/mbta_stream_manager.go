package stream

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type StreamManager struct {
	clients      map[chan string]struct{} // Map to store connected client channels
	clientsMutex sync.Mutex               // Mutex to safely update the clients map
	stop         chan struct{}            // Channel to signal when to stop streaming
}

// NewStreamManager initializes and returns a new StreamManager instance.
func NewStreamManager() *StreamManager {
	return &StreamManager{
		clients: make(map[chan string]struct{}),
		stop:    make(chan struct{}),
	}
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

// Broadcast sends the given data to all connected clients.
func (sm *StreamManager) Broadcast(data string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	for client := range sm.clients {
		select {
		case client <- data: // Send data to the client channel
		default: // Skip clients that are too slow to keep up
			log.Println("Client is slow, skipping...")
		}
	}
}

// StartStreaming connects to the MBTA API and continuously streams data to clients.
func (sm *StreamManager) StartStreaming(ctx context.Context, apiKey string) {
	go func() { // Run streaming logic in a separate goroutine
		url := "https://api-v3.mbta.com/vehicles?filter[route]=Red,Orange,Blue,Green-B,Green-C,Green-D,Green-E,Mattapan"
		// Create a new HTTP GET request with the given context
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return
		}
		// Set headers for SSE (Server-Sent Events) and API authentication
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("X-API-Key", apiKey)

		// Make the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to connect to MBTA API: %v", err)
			return
		}
		defer resp.Body.Close() // Ensure the response body is closed when done

		// Create a scanner to read the response line by line
		scanner := bufio.NewScanner(resp.Body)
		buffer := make([]byte, 1024*1024) // Set a 1 MB buffer size
		scanner.Buffer(buffer, len(buffer))

		var eventBuffer []string // Buffer to accumulate lines for a single SSE event
		for scanner.Scan() {
			select {
			case <-ctx.Done(): // Stop processing if the context is canceled
				return
			default:
				line := scanner.Text() // Read the next line from the stream

				// Check for empty lines, which signal the end of an event
				if line == "" {
					if len(eventBuffer) > 0 {
						fullEvent := strings.Join(eventBuffer, "\n") // Combine buffered lines
						sm.processSSE(fullEvent)                     // Process the complete event
						eventBuffer = []string{}                     // Clear the buffer
					}
					continue
				}

				// Accumulate lines for a single event
				eventBuffer = append(eventBuffer, line)
			}
		}

		// Handle errors that may occur while reading the stream
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stream: %v", err)
		}
	}()
}

// processSSE parses a complete SSE event and broadcasts it to clients.
func (sm *StreamManager) processSSE(event string) {
	lines := strings.Split(event, "\n") // Split the event into individual lines
	var eventType string                // Holds the "event" field
	var eventData []string              // Accumulates "data" fields

	// Parse each line of the event
	for _, line := range lines {
		if strings.HasPrefix(line, "event:") { // Extract the event type
			eventType = strings.TrimSpace(line[len("event:"):])
		} else if strings.HasPrefix(line, "data:") { // Extract data lines
			eventData = append(eventData, strings.TrimSpace(line[len("data:"):]))
		}
	}

	// Combine all data lines into a single payload
	fullData := strings.Join(eventData, "\n")

	// Format the event as SSE-compliant and broadcast it
	if fullData != "" {
		formattedEvent := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, fullData)
		sm.Broadcast(formattedEvent)
	}
}

// Stop stops the stream manager and signals all processes to stop.
func (sm *StreamManager) Stop() {
	close(sm.stop) // Close the stop channel to signal shutdown
}

// Global instance of StreamManager
var MBTAStreamManager = NewStreamManager()
