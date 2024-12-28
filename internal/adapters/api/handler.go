package api

import (
	"encoding/json"
	"github.com/yourusername/gtfs-explorer/internal/usecases"
	"net/http"
)

func FetchRoutes(fetchData usecases.FetchData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routes, err := fetchData.GetRoutes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	}
}

func FetchStops(fetchData usecases.FetchData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeID := r.URL.Query().Get("route_id")
		stops, err := fetchData.GetStops(routeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stops)
	}
}

func UpdateLiveData(fetchData usecases.FetchData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeID := r.URL.Query().Get("route_id")
		vehicles, err := fetchData.GetLiveData(routeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vehicles)
	}
}
