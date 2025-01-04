package models

type Stop struct {
	ID         string         `json:"id"`
	Attributes StopAttributes `json:"attributes"`
}

type StopAttributes struct {
	Address            *string `json:"address"`
	AtStreet           string  `json:"at_street"`
	Description        *string `json:"description"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Municipality       string  `json:"municipality"`
	Name               string  `json:"name"`
	OnStreet           string  `json:"on_street"`
	PlatformCode       *string `json:"platform_code"`
	PlatformName       *string `json:"platform_name"`
	VehicleType        int     `json:"vehicle_type"`
	WheelchairBoarding int     `json:"wheelchair_boarding"`
}

type StopsResponse struct {
	Data []Stop `json:"data"`
}
