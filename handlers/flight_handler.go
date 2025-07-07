package handlers

import (
	"airline-system/services"
	"encoding/json"
	"net/http"
	"time"
)

type FlightHandler struct {
	FlightService *services.FlightService
}

func NewFlightHandler(flightService *services.FlightService) *FlightHandler {
	return &FlightHandler{FlightService: flightService}
}

type AddFlightRequest struct {
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	DepartureTime string `json:"departure_time"` // ISO8601 string
	TotalSeats    int    `json:"total_seats"`
}

// POST /flights
func (h *FlightHandler) AddFlight(w http.ResponseWriter, r *http.Request) {
	var req AddFlightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Parse time
	depTime, err := time.Parse(time.RFC3339, req.DepartureTime)
	if err != nil {
		http.Error(w, "Invalid departure_time format, use RFC3339", http.StatusBadRequest)
		return
	}

	flight := h.FlightService.AddFlight(req.Source, req.Destination, depTime, req.TotalSeats)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(flight)
}

// GET /flights/search?source=abc&destination=xyz
func (h *FlightHandler) SearchFlights(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	source := query.Get("source")
	destination := query.Get("destination")

	if source == "" || destination == "" {
		http.Error(w, "Missing source or destination query params", http.StatusBadRequest)
		return
	}

	flights := h.FlightService.SearchFlights(source, destination)
	json.NewEncoder(w).Encode(flights)
}
