package models

type Stop struct {
	RouteID string  `json:"route_id"`
	StopID  string  `json:"stop_id"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}
