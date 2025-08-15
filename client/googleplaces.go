package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"tourist-api/model"
)

// Interface adhering to Interface Segregation Principle
type PlacesClient interface {
	GetTouristPlaces(lat, lon string, radius int) ([]model.Place, error)
}

// GooglePlacesClient implements PlacesClient
type GooglePlacesClient struct {
	APIKey     string
	HTTPClient *http.Client
}

type googlePlacesResponse struct {
	Results []struct {
		Name     string `json:"name"`
		Vicinity string `json:"vicinity"`
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

// Constructor for GooglePlacesClient (Single Responsibility)
func NewGooglePlacesClient(apiKey string) *GooglePlacesClient {
	return &GooglePlacesClient{
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *GooglePlacesClient) GetTouristPlaces(lat, lon string, radius int) ([]model.Place, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=%d&type=tourist_attraction&key=%s",
		lat, lon, radius, c.APIKey)
	
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var gr googlePlacesResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return nil, err
	}

	places := make([]model.Place, 0, len(gr.Results))
	for _, r := range gr.Results {
		places = append(places, model.Place{
			Name:     r.Name,
			Vicinity: r.Vicinity,
			Latitude: r.Geometry.Location.Lat,
			Longitude: r.Geometry.Location.Lng,
		})
	}

	return places, nil
}
