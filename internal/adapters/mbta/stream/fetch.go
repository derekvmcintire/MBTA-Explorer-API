package mbta

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// fetchStream handles the HTTP request logic to establish a stream connection
// to the specified URL and returns the response body for further processing.
//
// Parameters:
// - ctx: Context to manage request lifecycle and handle cancellations or timeouts.
// - url: The URL to connect to for the stream.
// - apiKey: API key for authentication.
//
// Returns:
// - io.ReadCloser: The response body for reading the stream data.
// - error: An error if the request fails or the response status is not OK.
func (m *MBTAStreamSource) fetchStream(ctx context.Context, url, apiKey string) (io.ReadCloser, error) {
	// Create a new HTTP GET request with the provided context, URL, and API key.
	req, err := m.createRequest(ctx, url, apiKey)
	if err != nil {
		return nil, err // Return error if the request creation fails
	}

	// Initialize an HTTP client and make the GET request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err // Return error if the request execution fails
	}

	// Verify that the response status code indicates success (200 OK).
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // Close the response body to avoid resource leaks
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Return the response body for reading the stream.
	return resp.Body, nil
}
