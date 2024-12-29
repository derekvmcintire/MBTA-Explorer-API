package usecases

import (
	"explorer/internal/domain/models" // Import the models package to work with the domain data structures
	"explorer/internal/ports/core"
	"explorer/internal/ports/data" // Import the data adapter package to interact with the external data layer
)

// fetchFromMBTAUseCaseImpl is a concrete implementation of the FetchFromMBTAUseCase interface
// It holds a reference to a client that interacts with the data layer (in this case, the MBTA API)
type fetchFromMBTAUseCaseImpl struct {
	client data.MBTAClient // The client used to fetch data from the MBTA API
}

// NewFetchData is a constructor function that creates and returns a new instance of fetchFromMBTAUseCaseImpl
// It accepts a data.MBTAClient as an argument to initialize the client for data fetching
func NewFetchData(client data.MBTAClient) core.FetchFromMBTAUseCase {
	return &fetchFromMBTAUseCaseImpl{client: client} // Return a new fetchFromMBTAUseCaseImpl instance with the provided client
}

// GetStops retrieves a list of stops for the given routeID
// It calls the FetchStops method on the client to fetch the data
func (f *fetchFromMBTAUseCaseImpl) GetStops(routeID string) ([]models.Stop, error) {
	return f.client.FetchStops(routeID)
}

// GetShapes retrieves a list of decoded coordinates for the given routeID
// It calls the FetchShapes method on the client to fetch the data
func (f *fetchFromMBTAUseCaseImpl) GetShapes(routeID string) (models.DecodedRouteShape, error) {
	return f.client.FetchShapes(routeID)
}

// GetLiveData retrieves live vehicle data for the given routeID
// It calls the FetchLiveData method on the client to fetch the data
func (f *fetchFromMBTAUseCaseImpl) GetLiveData(routeID string) ([]models.Vehicle, error) {
	return f.client.FetchLiveData(routeID)
}
