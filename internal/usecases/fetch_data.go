package usecases

import (
	"github.com/yourusername/gtfs-explorer/internal/adapters/data"
	"github.com/yourusername/gtfs-explorer/internal/domain/models"
)

type FetchData interface {
	GetRoutes() ([]models.Route, error)
	GetStops(routeID string) ([]models.Stop, error)
	GetLiveData(routeID string) ([]models.Vehicle, error)
}

type fetchData struct {
	client data.MBTAClient
}

func NewFetchData(client data.MBTAClient) FetchData {
	return &fetchData{client: client}
}

func (f *fetchData) GetRoutes() ([]models.Route, error) {
	return f.client.FetchRoutes()
}

func (f *fetchData) GetStops(routeID string) ([]models.Stop, error) {
	return f.client.FetchStops(routeID)
}

func (f *fetchData) GetLiveData(routeID string) ([]models.Vehicle, error) {
	return f.client.FetchLiveData(routeID)
}
