package user

import (
	"context"
	"errors"
	"sync"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/domain/user"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// InMemoryRepository implements user.Repository interface using in-memory storage
type InMemoryRepository struct {
	users map[string]*user.User
	mu    sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory user repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]*user.User),
	}
}

// Create adds a new user to the repository
func (r *InMemoryRepository) Create(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user with email already exists
	for _, existingUser := range r.users {
		if existingUser.Email == u.Email {
			return ErrUserAlreadyExists
		}
	}

	// Generate ID if not present
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	r.users[u.ID] = u
	return nil
}

// FindByEmail retrieves a user by email
func (r *InMemoryRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

// FindByID retrieves a user by ID
func (r *InMemoryRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return u, nil
}

// Update updates an existing user
func (r *InMemoryRepository) Update(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[u.ID]; !exists {
		return ErrUserNotFound
	}

	r.users[u.ID] = u
	return nil
}
