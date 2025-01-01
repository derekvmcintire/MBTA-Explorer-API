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
	clients      map[chan string]struct{}
	clientsMutex sync.Mutex
	stop         chan struct{}
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		clients: make(map[chan string]struct{}),
		stop:    make(chan struct{}),
	}
}

// AddClient adds a new client channel to receive data.
func (sm *StreamManager) AddClient(client chan string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	sm.clients[client] = struct{}{}
}

// RemoveClient removes a client channel when they disconnect.
func (sm *StreamManager) RemoveClient(client chan string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	if _, ok := sm.clients[client]; ok {
		delete(sm.clients, client)
		close(client)
	}
}

// Broadcast sends data to all connected clients.
func (sm *StreamManager) Broadcast(data string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	for client := range sm.clients {
		select {
		case client <- data:
		default: // Skip clients that can't keep up
			log.Println("Client is slow, skipping...")
		}
	}
}

// StartStreaming connects to the MBTA API and streams data to clients.
func (sm *StreamManager) StartStreaming(ctx context.Context, url, apiKey string) {
	go func() {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("X-API-Key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to connect to MBTA API: %v", err)
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		var eventBuffer []string
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
				line := scanner.Text()

				// Handle empty lines (event separator in SSE)
				if line == "" {
					if len(eventBuffer) > 0 {
						fullEvent := strings.Join(eventBuffer, "\n")
						sm.processSSE(fullEvent)
						eventBuffer = []string{} // Clear the buffer
					}
					continue
				}

				// Accumulate lines for a single event
				eventBuffer = append(eventBuffer, line)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stream: %v", err)
		}
	}()
}

// processSSE parses a full Server-Sent Event and broadcasts its data.
func (sm *StreamManager) processSSE(event string) {
	lines := strings.Split(event, "\n")
	var eventType string
	var eventData []string

	for _, line := range lines {
		if strings.HasPrefix(line, "event:") {
			eventType = strings.TrimSpace(line[len("event:"):])
		} else if strings.HasPrefix(line, "data:") {
			eventData = append(eventData, strings.TrimSpace(line[len("data:"):]))
		}
	}

	// Combine all data lines into a single payload
	fullData := strings.Join(eventData, "\n")

	// Log or broadcast the parsed event and data
	log.Printf("Event: %s, Data: %s", eventType, fullData)
	if fullData != "" {
		sm.Broadcast(fmt.Sprintf(`{"event": %q, "data": %s}`, eventType, fullData))
	}
}

// Stop stops the stream manager.
func (sm *StreamManager) Stop() {
	close(sm.stop)
}

var MBTAStreamManager = NewStreamManager()
