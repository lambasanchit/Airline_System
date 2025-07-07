package services

import (
	"airline-system/models"
	"airline-system/repository"
	"airline-system/utils"
	"strings"
	"time"
)

type FlightService struct {
	FlightRepo *repository.FlightRepository
}

func NewFlightService(flightRepo *repository.FlightRepository) *FlightService {
	return &FlightService{FlightRepo: flightRepo}
}

func (s *FlightService) AddFlight(source, destination string, departureTime time.Time, totalSeats int) *models.Flight {
	flight := &models.Flight{
		ID:             utils.GenerateID(),
		Source:         strings.ToLower(source),
		Destination:    strings.ToLower(destination),
		DepartureTime:  departureTime,
		TotalSeats:     totalSeats,
		AvailableSeats: totalSeats,
	}
	s.FlightRepo.Save(flight)
	return flight
}

func (s *FlightService) SearchFlights(source, destination string) []*models.Flight {
	source = strings.ToLower(source)
	destination = strings.ToLower(destination)

	var result []*models.Flight
	all := s.FlightRepo.FindAll()

	for _, f := range all {
		if f.Source == source && f.Destination == destination {
			result = append(result, f)
		}
	}
	return result
}

func (s *FlightService) GetFlightByID(id string) (*models.Flight, bool) {
	return s.FlightRepo.FindByID(id)
}
