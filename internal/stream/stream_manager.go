package stream

import (
	"context"
	"log"
)

// StartStreaming connects to the MBTA API and continuously streams data to clients.
func (sm *StreamManager) StartStreaming(ctx context.Context, apiKey string) {
	go func() {
		// Fetch the response body from the MBTA API
		respBody, err := sm.fetchStream(ctx, apiKey)
		if err != nil {
			log.Printf("Failed to fetch stream: %v", err)
			return
		}
		defer respBody.Close() // Ensure the response body is closed when done

		// Scan and process the stream
		sm.scanStream(ctx, respBody)
	}()
}
