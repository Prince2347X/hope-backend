package models

type HospitalModel struct {
	Id   	  string  `json:"id"`
	Name   	  string  `json:"name"`
	Longitude float64 `json:"location"`
	Latitude  float64 `json:"latitude"`
	Distance  float64 `json:"distance"`
}