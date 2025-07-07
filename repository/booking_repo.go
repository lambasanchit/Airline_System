package repository

import (
	"airline-system/models"
	"sync"
)

type BookingRepository struct {
	mu       sync.RWMutex
	bookings map[string]*models.Booking
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		bookings: make(map[string]*models.Booking),
	}
}

func (r *BookingRepository) Save(booking *models.Booking) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bookings[booking.ID] = booking
}

func (r *BookingRepository) FindByID(id string) (*models.Booking, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	booking, ok := r.bookings[id]
	return booking, ok
}

func (r *BookingRepository) FindAll() []*models.Booking {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []*models.Booking
	for _, booking := range r.bookings {
		list = append(list, booking)
	}
	return list
}
