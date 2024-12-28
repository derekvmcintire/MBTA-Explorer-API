package models

type Vehicle struct {
	ID    string  `json:"id"`
	Route string  `json:"route"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}
