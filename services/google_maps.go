package services

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Prince2347X/hope-backend/models"
)

type GoogleMapsServices struct{}

func (GoogleMapsServices) FetchHospitals(lat, long float64) ([]models.HospitalModel, error) {
	// MARK: TODO: - Fetching hospitals using the Maps API

	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")

	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=hospitals%20%26%20clinics&location=21.134633776764236%2C81.66728013054477&key=" + apiKey
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	fmt.Println("Response body:", string(body))

	return []models.HospitalModel{
		{Id: "1", Name: "Hospital 1", Location: models.Location{Lat: 0.0, Lng: 0.0}, Distance: 0.0},
	}, nil
}
