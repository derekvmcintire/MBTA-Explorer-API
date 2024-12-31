package usecases

import (
	"encoding/json"
	"explorer/internal/domain/models"
	"explorer/internal/ports/core"
	"explorer/internal/ports/data"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

type fetchFromMBTAUseCaseImpl struct {
	client data.MBTAClient  // The client used to fetch data from the MBTA API
	cache  *memcache.Client // Cache client for storing and retrieving data
}

// NewFetchData initializes fetchFromMBTAUseCaseImpl with a client and cache
func NewFetchData(client data.MBTAClient, cache *memcache.Client) core.FetchFromMBTAUseCase {
	return &fetchFromMBTAUseCaseImpl{client: client, cache: cache}
}

// GetStops retrieves a list of stops for the given routeID with caching
func (f *fetchFromMBTAUseCaseImpl) GetStops(routeID string) ([]models.Stop, error) {
	cacheKey := "stops:" + routeID
	item, err := f.cache.Get(cacheKey)
	if err == nil {
		log.Println("Cache hit for GetStops:", routeID)
		var stops []models.Stop
		if err := json.Unmarshal(item.Value, &stops); err == nil {
			return stops, nil
		}
		log.Println("Failed to unmarshal cached data, fetching fresh data.")
	}

	// Cache miss or unmarshalling failure
	stops, err := f.client.FetchStops(routeID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	value, _ := json.Marshal(stops)
	if err := f.cache.Set(&memcache.Item{Key: cacheKey, Value: value}); err != nil {
		log.Println("Failed to cache data for GetStops:", err)
	}

	return stops, nil
}

// GetShapes retrieves a list of decoded coordinates for the given routeID with caching
func (f *fetchFromMBTAUseCaseImpl) GetShapes(routeID string) (models.DecodedRouteShape, error) {
	cacheKey := "shapes:" + routeID
	item, err := f.cache.Get(cacheKey)
	if err == nil {
		log.Println("Cache hit for GetShapes:", routeID)
		var shapes models.DecodedRouteShape
		if err := json.Unmarshal(item.Value, &shapes); err == nil {
			return shapes, nil
		}
		log.Println("Failed to unmarshal cached data, fetching fresh data.")
	}

	// Cache miss or unmarshalling failure
	shapes, err := f.client.FetchShapes(routeID)
	if err != nil {
		return models.DecodedRouteShape{}, err
	}

	// Cache the result
	value, _ := json.Marshal(shapes)
	if err := f.cache.Set(&memcache.Item{Key: cacheKey, Value: value}); err != nil {
		log.Println("Failed to cache data for GetShapes:", err)
	}

	return shapes, nil
}

// GetLiveData retrieves live vehicle data for the given routeID without caching
func (f *fetchFromMBTAUseCaseImpl) GetLiveData(routeID string) ([]models.Vehicle, error) {
	return f.client.FetchLiveData(routeID)
}
