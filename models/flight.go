package models

import "time"

type Flight struct {
	ID             string    `json:"id"`
	Source         string    `json:"source"`
	Destination    string    `json:"destination"`
	DepartureTime  time.Time `json:"departure_time"`
	TotalSeats     int       `json:"total_seats"`
	AvailableSeats int       `json:"available_seats"`
}
