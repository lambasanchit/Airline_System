package services

import (
	"airline-system/models"
	"airline-system/repository"
	"airline-system/utils"
	"errors"
	"time"
)

type PaymentService struct {
	BookingRepo *repository.BookingRepository
	PaymentRepo *repository.PaymentRepository
}

func NewPaymentService(bookingRepo *repository.BookingRepository, paymentRepo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		BookingRepo: bookingRepo,
		PaymentRepo: paymentRepo,
	}
}

// Simulate payment process and update booking status
func (s *PaymentService) MakePayment(bookingID string, amount float64) (*models.Payment, error) {
	booking, ok := s.BookingRepo.FindByID(bookingID)
	if !ok {
		return nil, errors.New("booking not found")
	}

	if booking.Status != "Pending" {
		return nil, errors.New("payment already processed or cancelled")
	}

	// Simulate payment success
	payment := &models.Payment{
		ID:        utils.GenerateID(),
		BookingID: bookingID,
		Amount:    amount,
		Status:    "Success",
		CreatedAt: time.Now(),
	}

	// Save payment
	s.PaymentRepo.Save(payment)

	// Update booking status
	booking.Status = "Confirmed"

	return payment, nil
}
