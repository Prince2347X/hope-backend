package models

type HospitalModel struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	Location     Location `json:"location"`
	Distance     float64  `json:"distance"`
	Rating       float64  `json:"rating"`
	RatingsCount int      `json:"ratings_count"`
	IsOpenNow    bool     `json:"is_open_now"`
}
