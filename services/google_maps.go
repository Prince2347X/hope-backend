package services

import "github.com/Prince2347X/hope-backend/models"

type GoogleMapsServices struct{}

func (GoogleMapsServices) FetchHospitals(lat, long float64) []models.HospitalModel {
	// MARK: TODO: - Fetching hospitals using the Maps API

	return []models.HospitalModel{
		{Id: "1", Name: "Hospital 1", Location: models.Location{Lat: 0.0, Lng: 0.0}, Distance: 0.0},
	}
}