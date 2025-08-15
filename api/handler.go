package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tourist-api/client"
)

const{
	RADIUS = 2000
	LIMIT = 10
}

// Handler struct to respect Single Responsibility and Dependency Injection
type Handler struct {
	PlacesClient client.PlacesClient
}

func (h *Handler) TouristDestinationsHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	limitStr := r.URL.Query().Get("limit")

	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon parameter", http.StatusBadRequest)
		return
	}

	limit := LIMIT
	if limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil && val > 0 {
			limit = val
		}
	}

	radius := RADIUS // meters; this could be parameterized

	places, err := h.PlacesClient.GetTouristPlaces(lat, lon, radius)
	if err != nil {
		http.Error(w, "Failed to fetch places: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Limit results
	if len(places) > limit {
		places = places[:limit]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(places)
}
