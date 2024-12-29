package data

import (
	"explorer/internal/domain/models"
)

// MBTAClient is an interface that defines methods for fetching stops, shapes, and live vehicle data from the MBTA API
type MBTAClient interface {
	FetchStops(routeID string) ([]models.Stop, error)             // Method to fetch stops for a given route
	FetchShapes(routeID string) (models.DecodedRouteShape, error) // Method to fetch shapes for a given route
	FetchLiveData(routeID string) ([]models.Vehicle, error)       // Method to fetch live vehicle data for a given route
}
