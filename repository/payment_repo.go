package repository

import (
	"airline-system/models"
	"sync"
)

type PaymentRepository struct {
	mu       sync.RWMutex
	payments map[string]*models.Payment
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		payments: make(map[string]*models.Payment),
	}
}

func (r *PaymentRepository) Save(payment *models.Payment) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.payments[payment.ID] = payment
}

func (r *PaymentRepository) FindByID(id string) (*models.Payment, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	payment, ok := r.payments[id]
	return payment, ok
}
