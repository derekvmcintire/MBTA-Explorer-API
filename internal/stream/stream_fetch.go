package stream

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// fetchStream handles the HTTP request logic and returns the response body.
func (sm *StreamManager) fetchStream(ctx context.Context, apiKey string) (io.ReadCloser, error) {
	// Create a new HTTP GET request with the given context
	req, err := sm.createRequest(ctx, sm.url, apiKey)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check if the response status code indicates success
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // Close the body before returning an error
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body, nil
}
