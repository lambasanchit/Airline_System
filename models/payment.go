package models

import "time"

type Payment struct {
	ID        string    `json:"id"`
	BookingID string    `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
