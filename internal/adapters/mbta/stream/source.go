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

func NewMBTAStreamSource(distributor ports.StreamDistributor) *MBTAStreamSource {
	return &MBTAStreamSource{
		distributor: distributor,
	}
}

// createRequest creates an HTTP request for streaming data from the MBTA API.
// It sets the appropriate headers for the request, including the API key and content type.
func (m *MBTAStreamSource) createRequest(ctx context.Context, url, apiKey string) (*http.Request, error) {
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

// stream/utils.go - updated Start method
func (m *MBTAStreamSource) Start(ctx context.Context, url, apiKey string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Context cancelled, stopping stream")
				return
			default:
				respBody, err := m.fetchStream(ctx, url, apiKey)
				if err != nil {
					log.Printf("Failed to fetch stream: %v", err)
					// Add a small delay before retrying to prevent tight loops on persistent errors
					time.Sleep(time.Second * 5)
					continue
				}

				// Start stream processing in a separate goroutine
				processDone := make(chan struct{})
				go func() {
					defer close(processDone)
					defer respBody.Close()
					m.scanStream(ctx, respBody)
				}()

				// Wait for either context cancellation or stream processing to finish
				select {
				case <-ctx.Done():
					respBody.Close()
					return
				case <-processDone:
					log.Println("Stream processing ended, will retry")
					// Continue the outer loop to restart the stream
				}
			}
		}
	}()
}
