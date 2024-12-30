package data

import (
	"encoding/json"
	"explorer/internal/domain/models"
	"explorer/internal/ports/data"
	"explorer/internal/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

const mbtaAPIBaseUrl = "https://api-v3.mbta.com"

// mbtaClientImpl is the implementation of the MBTAClient interface
// It holds the API key and HTTP client used for making requests
type mbtaClientImpl struct {
	apiKey string
	client *http.Client
}

// NewMBTAClient is a constructor function that initializes and returns a new instance of mbtaClientImpl
func NewMBTAClient(apiKey string) data.MBTAClient {
	return &mbtaClientImpl{
		apiKey: apiKey,                                  // Set the API key from the argument
		client: &http.Client{Timeout: 10 * time.Second}, // Set a timeout of 10 seconds for HTTP requests
	}
}

// FetchShapes fetches the shape data for a given route ID from the MBTA API
func (m *mbtaClientImpl) FetchShapes(routeID string) (models.DecodedRouteShape, error) {
	// Format the endpoint URL to include the route ID in the query parameters
	endpoint := fmt.Sprintf("%s/shapes?filter[route]=%s", mbtaAPIBaseUrl, routeID)

	// Call fetchData to get the raw data from the API
	data, err := m.fetchData(endpoint)
	if err != nil {
		return models.DecodedRouteShape{}, fmt.Errorf("failed to fetch shapes: %w", err)
	}

	// Declare a variable to hold the response data
	var shapeResponse models.ShapeResponse

	// Unmarshal the raw JSON data into the shapeResponse variable
	if err := json.Unmarshal(data, &shapeResponse); err != nil {
		log.Fatal(err) // Log a fatal error if unmarshalling fails (this will stop execution)
	}

	// Decode the shape data into coordinates
	var decodedRouteShape models.DecodedRouteShape
	decodedRouteShape.RouteID = routeID
	decodedRouteShape.Coordinates, err = utils.DecodeShapes(shapeResponse.Data)
	if err != nil {
		return models.DecodedRouteShape{}, fmt.Errorf("failed to decode shapes: %w", err)
	}

	// Return the decoded route shape
	return decodedRouteShape, nil
}

// FetchStops fetches the list of stops for a given route ID from the MBTA API
func (m *mbtaClientImpl) FetchStops(routeID string) ([]models.Stop, error) {
	// Format the endpoint URL to include the route ID in the query parameters
	endpoint := fmt.Sprintf("%s/stops?filter[route]=%s", mbtaAPIBaseUrl, routeID)

	// Call fetchData to get the raw data from the API
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err // Return the error if fetching the data fails
	}

	// Declare a variable to hold the response data (which is a list of stops)
	var stopsResponse models.StopsResponse

	// Unmarshal the raw JSON data into the stopsResponse variable
	if err := json.Unmarshal(data, &stopsResponse); err != nil {
		log.Fatal(err) // Log a fatal error if unmarshalling fails (this will stop execution)
	}

	// Return the list of stops from the response data
	return stopsResponse.Data, nil
}

// FetchLiveData fetches the live vehicle data for a given route ID from the MBTA API
func (m *mbtaClientImpl) FetchLiveData(routeID string) ([]models.Vehicle, error) {
	// Format the endpoint URL to include the route ID in the query parameters
	endpoint := fmt.Sprintf("%s/vehicles?filter[route]=%s", mbtaAPIBaseUrl, routeID)

	// Call fetchData to get the raw data from the API
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err // Return the error if fetching the data fails
	}

	// Declare a variable to hold the list of vehicles
	var vehicles models.VehicleResponse

	// Unmarshal the raw JSON data into the vehicles slice
	if err := json.Unmarshal(data, &vehicles); err != nil {
		return nil, err // Return the error if unmarshalling fails
	}

	// Return the list of vehicles
	return vehicles.Data, nil
}
