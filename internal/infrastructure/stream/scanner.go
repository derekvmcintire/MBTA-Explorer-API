package stream

import (
	"bufio"
	"context"
	"io"
	"log"
	"strings"
)

// scanStream processes the response body stream and handles server-sent events (SSE).
func (sm *StreamManager) scanStream(ctx context.Context, responseBody io.Reader) {
	log.Println("scanStream has been called")
	scanner := bufio.NewScanner(responseBody)
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
}
