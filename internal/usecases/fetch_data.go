package usecases

import (
	"explorer/internal/adapters/data" // Import the data adapter package to interact with the external data layer
	"explorer/internal/domain/models" // Import the models package to work with the domain data structures
)

// FetchData is an interface that defines the methods for fetching stops and live vehicle data
type FetchData interface {
	// GetStops fetches a list of stops for a given route ID
	GetStops(routeID string) ([]models.Stop, error)

	// GetLiveData fetches live vehicle data for a given route ID
	GetLiveData(routeID string) ([]models.Vehicle, error)
}

// fetchData is a concrete implementation of the FetchData interface
// It holds a reference to a client that interacts with the data layer (in this case, the MBTA API)
type fetchData struct {
	client data.MBTAClient // The client used to fetch data from the MBTA API
}

// NewFetchData is a constructor function that creates and returns a new instance of fetchData
// It accepts a data.MBTAClient as an argument to initialize the client for data fetching
func NewFetchData(client data.MBTAClient) FetchData {
	return &fetchData{client: client} // Return a new fetchData instance with the provided client
}

// GetStops retrieves a list of stops for the given routeID
// It calls the FetchStops method on the client to fetch the data
func (f *fetchData) GetStops(routeID string) ([]models.Stop, error) {
	// Call FetchStops on the client and return the result or any error
	return f.client.FetchStops(routeID)
}

// GetLiveData retrieves live vehicle data for the given routeID
// It calls the FetchLiveData method on the client to fetch the data
func (f *fetchData) GetLiveData(routeID string) ([]models.Vehicle, error) {
	// Call FetchLiveData on the client and return the result or any error
	return f.client.FetchLiveData(routeID)
}
