package data

import (
	"encoding/json"                   // Import the json package to handle JSON encoding and decoding
	"explorer/internal/domain/models" // Import the models package to work with the application data structures
	"fmt"                             // Import the fmt package for formatted I/O operations like string formatting
	"io"                              // Import the io package for handling I/O operations like reading from a body
	"log"                             // Import the log package for logging important information
	"net/http"                        // Import the net/http package for handling HTTP requests and responses
	"time"                            // Import the time package to manage timeouts and delays
)

// Define the base URL for the MBTA API
const mbtaAPIBase = "https://api-v3.mbta.com"

// MBTAClient is an interface that defines methods for fetching stops and live vehicle data from the MBTA API
type MBTAClient interface {
	FetchStops(routeID string) ([]models.Stop, error) // Method to fetch stops for a given route
	FetchShapes(routeID string) ([]models.Shape, error)
	FetchLiveData(routeID string) ([]models.Vehicle, error) // Method to fetch live vehicle data for a given route
}

// mbtaClientImpl is the implementation of the MBTAClient interface
// It holds the API key and HTTP client used for making requests
type mbtaClientImpl struct {
	apiKey string       // API key used for authentication when making requests to the MBTA API
	client *http.Client // HTTP client used for making requests with a timeout setting
}

// NewMBTAClient is a constructor function that initializes and returns a new instance of mbtaClientImpl
func NewMBTAClient(apiKey string) MBTAClient {
	return &mbtaClientImpl{
		apiKey: apiKey,                                  // Set the API key from the argument
		client: &http.Client{Timeout: 10 * time.Second}, // Set a timeout of 10 seconds for HTTP requests
	}
}

// fetchData is a helper method that makes a GET request to the given endpoint
// It returns the raw response body as a byte slice or an error if something goes wrong
func (m *mbtaClientImpl) fetchData(endpoint string) ([]byte, error) {

	// Create a new GET request with the given endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("Error building the request?") // Log if there's an error creating the request
		return nil, err
	}

	// Set the API key in the request header for authentication
	req.Header.Set("x-api-key", m.apiKey)

	// Execute the request using the HTTP client
	resp, err := m.client.Do(req)
	if err != nil {
		log.Println("error after the response is sent") // Log if there's an error executing the request
		return nil, err
	}
	defer resp.Body.Close() // Ensure the response body is closed after use

	// Check if the response status code is not OK (200)
	if resp.StatusCode != http.StatusOK {
		log.Println("status code not ok") // Log if the status code is not OK
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

func (m *mbtaClientImpl) FetchShapes(routeID string) ([]models.Shape, error) {
	endpoint := fmt.Sprintf("%s/shapes?filter[route]=%s", mbtaAPIBase, routeID)
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err
	}
	var shapeResponse models.ShapeResponse
	if err := json.Unmarshal([]byte(data), &shapeResponse); err != nil {
		log.Fatal(err) // Log a fatal error if unmarshalling fails (this will stop execution)
	}
	return shapeResponse.Data, nil
}

// FetchStops fetches the list of stops for a given route ID from the MBTA API
func (m *mbtaClientImpl) FetchStops(routeID string) ([]models.Stop, error) {
	// Format the endpoint URL to include the route ID in the query parameters
	endpoint := fmt.Sprintf("%s/stops?filter[route]=%s", mbtaAPIBase, routeID)

	// Call fetchData to get the raw data from the API
	data, err := m.fetchData(endpoint)

	// If there's an error fetching the data, return the error
	if err != nil {
		return nil, err
	}

	// Declare a variable to hold the response data (which is a list of stops)
	var stopsResponse models.StopsResponse

	// Unmarshal the raw JSON data into the stopsResponse variable
	if err := json.Unmarshal([]byte(data), &stopsResponse); err != nil {
		log.Fatal(err) // Log a fatal error if unmarshalling fails (this will stop execution)
	}

	// Return the list of stops from the response data
	return stopsResponse.Data, nil
}

// FetchLiveData fetches the live vehicle data for a given route ID from the MBTA API
func (m *mbtaClientImpl) FetchLiveData(routeID string) ([]models.Vehicle, error) {
	// Format the endpoint URL to include the route ID in the query parameters
	endpoint := fmt.Sprintf("%s/vehicles?filter[route]=%s", mbtaAPIBase, routeID)

	// Call fetchData to get the raw data from the API
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err // Return the error if fetching the data fails
	}

	// Declare a variable to hold the list of vehicles
	var vehicles []models.Vehicle

	// Unmarshal the raw JSON data into the vehicles slice
	if err := json.Unmarshal(data, &vehicles); err != nil {
		return nil, err // Return the error if unmarshalling fails
	}

	// Return the list of vehicles
	return vehicles, nil
}
