package core

import (
	"explorer/internal/domain/models"
)

// FetchData is an interface that defines the methods for fetching stops and live vehicle data
type FetchFromMBTAUseCase interface {
	// GetStops fetches a list of stops for a given route ID
	GetStops(routeID string) ([]models.Stop, error)

	// GetShapes fetches a list of decoded coordinates for a given route ID
	GetShapes(routeID string) (models.DecodedRouteShape, error)

	// GetLiveData fetches live vehicle data for a given route ID
	GetLiveData(routeID string) ([]models.Vehicle, error)
}
