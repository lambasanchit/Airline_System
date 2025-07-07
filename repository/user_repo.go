package repository

import (
	"airline-system/models"
	"errors"
	"sync"
)

var ErrUserExists = errors.New("user already exists")

type UserRepository struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*models.User),
	}
}

// Save new user
func (r *UserRepository) Save(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user email already exists (simple check)
	for _, u := range r.users {
		if u.Email == user.Email {
			return ErrUserExists
		}
	}

	r.users[user.ID] = user
	return nil
}

// Find user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, true
		}
	}
	return nil, false
}

// Find user by ID
func (r *UserRepository) FindByID(id string) (*models.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	return u, ok
}
