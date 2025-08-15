package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"tourist-api/api"
	"tourist-api/client"
)

func main() {
	apiKey := os.Getenv("GOOGLE_PLACES_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_PLACES_API_KEY environment variable is required")
	}

	placesClient := client.NewGooglePlacesClient(apiKey)
	handler := &api.Handler{PlacesClient: placesClient}

	http.HandleFunc("/api/tourist-destinations", handler.TouristDestinationsHandler)

	port := ":8080"
	fmt.Println("Server running on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
