package services

import (
	"airline-system/models"
	"airline-system/repository"
	"airline-system/utils"
	"errors"
	"sync"
	"time"
)

type BookingService struct {
	BookingRepo *repository.BookingRepository
	FlightRepo  *repository.FlightRepository
	lock        sync.Mutex
}

func NewBookingService(bookingRepo *repository.BookingRepository, flightRepo *repository.FlightRepository) *BookingService {
	return &BookingService{
		BookingRepo: bookingRepo,
		FlightRepo:  flightRepo,
	}
}

// Create a new booking for a user on a flight
func (s *BookingService) CreateBooking(userID, flightID string) (*models.Booking, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	flight, ok := s.FlightRepo.FindByID(flightID)
	if !ok {
		return nil, errors.New("flight not found")
	}

	if flight.AvailableSeats <= 0 {
		return nil, errors.New("no seats available")
	}

	// Book one seat
	seatNo := flight.TotalSeats - flight.AvailableSeats + 1
	flight.AvailableSeats--

	booking := &models.Booking{
		ID:        utils.GenerateID(),
		UserID:    userID,
		FlightID:  flightID,
		SeatNo:    seatNo,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	s.BookingRepo.Save(booking)
	return booking, nil
}

func (s *BookingService) CancelBooking(bookingID string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	booking, ok := s.BookingRepo.FindByID(bookingID)
	if !ok {
		return errors.New("booking not found")
	}

	if booking.Status == "Cancelled" {
		return errors.New("already cancelled")
	}

	// Free up the seat
	flight, ok := s.FlightRepo.FindByID(booking.FlightID)
	if ok {
		flight.AvailableSeats++
	}

	booking.Status = "Cancelled"
	return nil
}
