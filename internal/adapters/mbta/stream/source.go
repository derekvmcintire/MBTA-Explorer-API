package mbta

import (
	"context"
	ports "explorer/internal/ports/streaming"
	"log"
	"net/http"
	"time"
)

type MBTAStreamSource struct {
	distributor ports.StreamDistributor
}

// NewMBTAStreamSource initializes a new MBTAStreamSource with the given distributor.
func NewMBTAStreamSource(distributor ports.StreamDistributor) *MBTAStreamSource {
	return &MBTAStreamSource{
		distributor: distributor,
	}
}

// createRequest creates an HTTP GET request for streaming data from the MBTA API.
//
// Parameters:
// - ctx: The context to manage request lifecycle (e.g., timeouts, cancellations).
// - url: The endpoint to connect to.
// - apiKey: The API key for authorization.
//
// Returns:
// - A pointer to the created HTTP request or an error if the request creation fails.
func (m *MBTAStreamSource) createRequest(ctx context.Context, url, apiKey string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil) // Create a GET request.
	if err != nil {
		return nil, err // Return error if request creation fails.
	}
	// Set necessary headers for SSE.
	req.Header.Set("Accept", "text/event-stream") // Specify content type for SSE.
	req.Header.Set("X-API-Key", apiKey)           // Add API key for authorization.
	return req, nil
}

// Start begins streaming data from the MBTA API.
//
// Parameters:
// - ctx: The context to manage the streaming lifecycle (e.g., cancellations, timeouts).
// - url: The endpoint to fetch SSE data from.
// - apiKey: The API key for authenticating the request.
//
// This method:
// - Continuously attempts to fetch and process the stream unless the context is cancelled.
// - Implements retries with a delay upon errors to avoid tight retry loops.
func (m *MBTAStreamSource) Start(ctx context.Context, url, apiKey string) {
	go func() { // Run the streaming logic in a goroutine.
		for { // Infinite loop to retry on errors or disconnections.
			select {
			case <-ctx.Done(): // Exit loop if context is cancelled.
				log.Println("Context cancelled, stopping stream")
				return
			default:
				// Fetch the stream from the MBTA API.
				respBody, err := m.fetchStream(ctx, url, apiKey)
				if err != nil {
					log.Printf("Failed to fetch stream: %v", err)
					time.Sleep(time.Second * 5) // Delay before retrying to avoid tight loops.
					continue
				}

				// Process the stream in a separate goroutine.
				processDone := make(chan struct{}) // Channel to signal completion of stream processing.
				go func() {
					defer close(processDone)    // Ensure channel closure when processing finishes.
					defer respBody.Close()      // Ensure response body is closed.
					m.scanStream(ctx, respBody) // Scan and process the stream.
				}()

				// Wait for either context cancellation or stream processing to finish.
				select {
				case <-ctx.Done(): // Stop processing if context is cancelled.
					respBody.Close()
					return
				case <-processDone: // Restart the loop on stream processing completion.
					log.Println("Stream processing ended, will retry")
				}
			}
		}
	}()
}
