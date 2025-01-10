package mbta

import (
	"bufio"
	"context"
	"io"
	"log"
	"strings"
)

// scanStream reads and processes server-sent events (SSE) from the response body stream.
//
// Parameters:
// - ctx: The context to manage cancellation or timeouts.
// - responseBody: The stream to be read, typically the HTTP response body.
//
// Functionality:
// - Reads lines from the stream using a buffered scanner.
// - Buffers lines for each SSE message until a blank line indicates the end of the event.
// - Processes complete SSE messages and handles errors in the stream.
func (m *MBTAStreamSource) scanStream(ctx context.Context, responseBody io.Reader) {
	// Create a buffered scanner to read the response body line by line.
	scanner := bufio.NewScanner(responseBody)

	// Set a 1 MB buffer size for the scanner to handle large event streams.
	buffer := make([]byte, 1024*1024)
	scanner.Buffer(buffer, len(buffer))

	var eventBuffer []string // Buffer to accumulate lines for a single SSE event.

	// Loop through each line in the stream.
	for scanner.Scan() {
		select {
		case <-ctx.Done(): // Exit if the context is canceled.
			return
		default:
			line := scanner.Text() // Get the current line from the stream.

			// Check for an empty line signaling the end of an SSE event.
			if line == "" {
				// If the buffer has accumulated lines, process the event.
				if len(eventBuffer) > 0 {
					fullEvent := strings.Join(eventBuffer, "\n") // Combine buffered lines.
					m.processSSE(fullEvent)                      // Process the complete SSE event.
					eventBuffer = []string{}                     // Clear the buffer for the next event.
				}
				continue // Skip to the next line.
			}

			// Accumulate lines for the current SSE event.
			eventBuffer = append(eventBuffer, line)
		}
	}

	// Handle errors that may occur while scanning the stream.
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading stream: %v", err) // Log the error for debugging.
	}
}
