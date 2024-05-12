package models

type HospitalModel struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Distance float64  `json:"distance"`
}
