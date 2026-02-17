package user

import (
	"errors"
	"time"
)

var (
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrEmptyPasswordHash  = errors.New("password hash cannot be empty")
	ErrEmptyUsername      = errors.New("username cannot be empty")
)

// User represents the user domain entity
type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new user entity with validation
func NewUser(username, email, passwordHash string) (*User, error) {
	if username == "" {
		return nil, ErrEmptyUsername
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if passwordHash == "" {
		return nil, ErrEmptyPasswordHash
	}

	now := time.Now()
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
