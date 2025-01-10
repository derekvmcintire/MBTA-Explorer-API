package mbta

import (
	"fmt"
	"strings"
)

// processSSE parses a complete Server-Sent Events (SSE) message and broadcasts it to all connected clients.
// This function processes the raw event string, extracts relevant fields, and formats them
// into an SSE-compliant message before sending it to the Broadcast function.
func (m *MBTAStreamSource) processSSE(event string) {
	// Split the event string into individual lines to process each line separately.
	lines := strings.Split(event, "\n")

	var eventType string   // Holds the value of the "event" field in the SSE message.
	var eventData []string // Accumulates the "data" lines for the SSE message.

	// Iterate over each line in the event to extract the event type and data lines.
	for _, line := range lines {
		if strings.HasPrefix(line, "event:") {
			// If the line starts with "event:", extract and trim the event type value.
			eventType = strings.TrimSpace(line[len("event:"):])
		} else if strings.HasPrefix(line, "data:") {
			// If the line starts with "data:", extract and trim the data value.
			// Append the trimmed value to the eventData slice.
			eventData = append(eventData, strings.TrimSpace(line[len("data:"):]))
		}
	}

	// Combine all accumulated data lines into a single string, separated by newline characters.
	fullData := strings.Join(eventData, "\n")

	// If there is data to send, format it as an SSE-compliant message and broadcast it.
	if fullData != "" {
		// Format the SSE message with the extracted event type and data.
		formattedEvent := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, fullData)

		// Send the formatted event to all connected clients via the Broadcast method.
		m.distributor.Broadcast(formattedEvent)
	}
}
