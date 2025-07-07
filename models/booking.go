package models

import "time"

type Booking struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	FlightID  string    `json:"flight_id"`
	SeatNo    int       `json:"seat_no"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
