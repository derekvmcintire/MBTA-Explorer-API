package data

import (
	"encoding/json"
	"fmt"
	"github.com/yourusername/gtfs-explorer/internal/domain/models"
	"net/http"
	"time"
)

const mbtaAPIBase = "https://api-v3.mbta.com"

type MBTAClient interface {
	FetchRoutes() ([]models.Route, error)
	FetchStops(routeID string) ([]models.Stop, error)
	FetchLiveData(routeID string) ([]models.Vehicle, error)
}

type mbtaAPIClient struct {
	apiKey string
	client *http.Client
}

func NewMBTAClient(apiKey string) MBTAClient {
	return &mbtaAPIClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (m *mbtaAPIClient) fetchData(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", m.apiKey)
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %v", resp.Status)
	}

	return json.Marshal(resp.Body)
}

func (m *mbtaAPIClient) FetchRoutes() ([]models.Route, error) {
	endpoint := fmt.Sprintf("%s/routes", mbtaAPIBase)
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err
	}

	var routes []models.Route
	if err := json.Unmarshal(data, &routes); err != nil {
		return nil, err
	}

	return routes, nil
}

func (m *mbtaAPIClient) FetchStops(routeID string) ([]models.Stop, error) {
	endpoint := fmt.Sprintf("%s/stops?filter[route]=%s", mbtaAPIBase, routeID)
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err
	}

	var stops []models.Stop
	if err := json.Unmarshal(data, &stops); err != nil {
		return nil, err
	}

	return stops, nil
}

func (m *mbtaAPIClient) FetchLiveData(routeID string) ([]models.Vehicle, error) {
	endpoint := fmt.Sprintf("%s/vehicles?filter[route]=%s", mbtaAPIBase, routeID)
	data, err := m.fetchData(endpoint)
	if err != nil {
		return nil, err
	}

	var vehicles []models.Vehicle
	if err := json.Unmarshal(data, &vehicles); err != nil {
		return nil, err
	}

	return vehicles, nil
}
