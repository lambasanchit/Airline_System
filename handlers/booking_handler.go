package handlers

import (
	"airline-system/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookingHandler struct {
	BookingService *services.BookingService
}

func NewBookingHandler(bookingService *services.BookingService) *BookingHandler {
	return &BookingHandler{BookingService: bookingService}
}

type CreateBookingRequest struct {
	UserID   string `json:"user_id"`
	FlightID string `json:"flight_id"`
}

// POST /bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	booking, err := h.BookingService.CreateBooking(req.UserID, req.FlightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// POST /bookings/{id}/cancel
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.BookingService.CancelBooking(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking cancelled"})
}
