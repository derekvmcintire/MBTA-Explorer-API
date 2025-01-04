package stream

import (
	"context"
	"log"
	"sync"
)

type StreamManager struct {
	clients      map[chan string]struct{} // Map to store connected client channels
	clientsMutex sync.Mutex               // Mutex to safely update the clients map
	stop         chan struct{}            // Channel to signal when to stop streaming
}

// StartStreaming connects to the MBTA API and continuously streams data to clients.
func (sm *StreamManager) StartStreaming(ctx context.Context, url, apiKey string) {
	go func() {
		// Fetch the response body from the MBTA API
		respBody, err := sm.fetchStream(ctx, url, apiKey)
		if err != nil {
			log.Printf("Failed to fetch stream: %v", err)
			return
		}
		defer respBody.Close() // Ensure the response body is closed when done

		// Scan and process the stream
		sm.scanStream(ctx, respBody)
	}()
}
