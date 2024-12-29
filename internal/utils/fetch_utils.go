package utils

import (
	"fmt"      // Importing the fmt package for formatted I/O
	"io"       // Importing io package to handle reading from streams (response body)
	"net/http" // Importing the net/http package to make HTTP requests
	"time"     // Importing the time package for timing operations (e.g., retries with delay)
)

// FetchDataWithRetry attempts to fetch data from the specified URL.
// It will retry the request up to maxRetries times if an error occurs or the status code is not OK (200).
func FetchDataWithRetry(url string, maxRetries int) ([]byte, error) {
	// Iterate through the retry loop, trying up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Send the GET request to the URL
		resp, err := http.Get(url)

		// If the request is successful and the status code is 200 OK, process the response
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close() // Ensure the response body is closed after reading

			// Read the entire response body into a byte slice
			return io.ReadAll(resp.Body)
		}

		// If the request fails (either an error or non-200 status code), wait for 2 seconds before retrying
		time.Sleep(2 * time.Second)
	}

	// If the request failed after maxRetries attempts, return an error message
	return nil, fmt.Errorf("failed to fetch data after %d retries", maxRetries)
}
