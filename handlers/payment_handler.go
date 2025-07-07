package handlers

import (
	"airline-system/services"
	"encoding/json"
	"net/http"
)

type PaymentHandler struct {
	PaymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentService: paymentService}
}

type MakePaymentRequest struct {
	BookingID string  `json:"booking_id"`
	Amount    float64 `json:"amount"`
}

// POST /payment/pay
func (h *PaymentHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req MakePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	payment, err := h.PaymentService.MakePayment(req.BookingID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)
}
