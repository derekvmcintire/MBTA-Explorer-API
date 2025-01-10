package response

import "explorer/internal/core/domain/models"

type GetRouteResponse struct {
	ID          string        `json:"id"`
	Coordinates [][][]float64 `json:"coordinates"`
	Stops       []models.Stop `json:"stops"`
}
