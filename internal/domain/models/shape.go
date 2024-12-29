package models

type Shape struct {
	ID         string          `json:"id"`
	Attributes ShapeAttributes `json:"attributes"`
	Type       string          `json:"type"`
}

type ShapeAttributes struct {
	PolyLine string `json:"polyline"`
}

type ShapeResponse struct {
	Data []Shape `json:"data"`
}
