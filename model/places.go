package model

type Place struct {
	Name     string  `json:"name"`
	Vicinity string  `json:"vicinity"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
