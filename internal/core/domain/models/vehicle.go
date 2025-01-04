package models

type Vehicle struct {
	ID         string            `json:"id"`
	Route      string            `json:"route"`
	Attributes VehicleAttributes `json:"attributes"`
}

type VehicleAttributes struct {
	Bearing             int                `json:"bearing"`
	Carriages           []VehicleCarriages `json:"carriages"`
	CurrentStatus       string             `json:"current_status"`
	CurrentStopSequence int                `json:"current_stop_sequence"`
	Direction           int                `json:"direction"`
	Label               string             `json:"label"`
	Latitude            float64            `json:"latitude"`
	Longitude           float64            `json:"longitude"`
	OccupancyStatus     string             `json:"occupancy_status"`
	Revenue             string             `json:"revenue"`
	Speed               float64            `json:"speed"`
	UpdatedAt           string             `json:"updated_at"`
}

type VehicleCarriages struct {
	OccupancyStatus     string `json:"occupancy_status"`
	OccupancyPercentage int    `json:"occupancy_percentage"`
	Label               string `json:"label"`
}

type VehicleResponse struct {
	Data []Vehicle `json:"data"`
}
