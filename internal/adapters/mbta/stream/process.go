package mbta

import (
	"fmt"
	"strings"
)

// processSSE parses a Server-Sent Events (SSE) message, extracts its fields,
// formats it into an SSE-compliant message, and broadcasts it to connected clients.
//
// Parameters:
// - event: The raw SSE event string received from the server.
//
// Functionality:
// - Splits the raw event string into lines to parse individual fields.
// - Extracts the "event" and "data" fields from the message.
// - Formats the parsed fields into an SSE-compliant message.
// - Broadcasts the formatted message to all connected clients via the distributor.
func (m *MBTAStreamSource) processSSE(event string) {
	// Split the raw event string into lines for processing.
	lines := strings.Split(event, "\n")

	var eventType string   // Holds the extracted "event" field value.
	var eventData []string // Accumulates "data" field values.

	// Process each line to extract relevant SSE fields.
	for _, line := range lines {
		if strings.HasPrefix(line, "event:") {
			// Extract and trim the value of the "event" field.
			eventType = strings.TrimSpace(line[len("event:"):])
		} else if strings.HasPrefix(line, "data:") {
			// Extract and trim the value of the "data" field and add it to eventData.
			eventData = append(eventData, strings.TrimSpace(line[len("data:"):]))
		}
	}

	// Combine all data lines into a single string, separated by newline characters.
	fullData := strings.Join(eventData, "\n")

	// If data exists, format and broadcast the SSE-compliant message.
	if fullData != "" {
		// Format the SSE message with the event type and data.
		formattedEvent := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, fullData)

		// Broadcast the formatted message to all connected clients.
		m.distributor.Broadcast(formattedEvent)
	}
}
