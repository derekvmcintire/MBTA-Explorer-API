package data

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// fetchData is a helper method that makes a GET request to the given endpoint
// It returns the raw response body as a byte slice or an error if something goes wrong
func (m *mbtaClientImpl) fetchData(endpoint string) ([]byte, error) {
	// Create a new GET request with the given endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("Error building the request") // Log if there's an error creating the request
		return nil, err
	}

	// Set the API key in the request header for authentication
	req.Header.Set("x-api-key", m.apiKey)

	// Execute the request using the HTTP client
	resp, err := m.client.Do(req)
	if err != nil {
		log.Println("Error after the response is sent") // Log if there's an error executing the request
		return nil, err
	}
	defer resp.Body.Close() // Ensure the response body is closed after use

	// Check if the response status code is not OK (200)
	if resp.StatusCode != http.StatusOK {
		log.Println("Status code not OK") // Log if the status code is not OK
		return nil, fmt.Errorf("error fetching data: %v", resp.Status)
	}

	// Read the entire response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading the response body") // Log if there's an error reading the body
		return nil, err
	}

	// Return the response body as raw data
	return body, nil
}
