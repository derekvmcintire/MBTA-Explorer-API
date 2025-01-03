package stream

import (
	"fmt"
	"strings"
)

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
