package repository

import (
	"airline-system/models"
	"sync"
)

type FlightRepository struct {
	mu      sync.RWMutex
	flights map[string]*models.Flight
}

func NewFlightRepository() *FlightRepository {
	return &FlightRepository{
		flights: make(map[string]*models.Flight),
	}
}

func (r *FlightRepository) Save(flight *models.Flight) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.flights[flight.ID] = flight
}

func (r *FlightRepository) FindAll() []*models.Flight {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []*models.Flight
	for _, flight := range r.flights {
		list = append(list, flight)
	}
	return list
}

func (r *FlightRepository) FindByID(id string) (*models.Flight, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	flight, ok := r.flights[id]
	return flight, ok
}
