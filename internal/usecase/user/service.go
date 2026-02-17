package user

import (
	"context"
	"errors"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/domain/user"
	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
	"github.com/DiaaSaada/crypto-payment-gateway/pkg/password"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UseCase defines the interface for user business logic
type UseCase interface {
	Register(ctx context.Context, username, email, pwd string) (*user.User, error)
	Login(ctx context.Context, email, pwd string) (string, error)
}

// Service implements UseCase interface
type Service struct {
	repo       user.Repository
	jwtService *jwt.Service
}

// NewService creates a new user service
func NewService(repo user.Repository, jwtService *jwt.Service) *Service {
	return &Service{
		repo:       repo,
		jwtService: jwtService,
	}
}

// Register creates a new user account
func (s *Service) Register(ctx context.Context, username, email, pwd string) (*user.User, error) {
	// Check if user already exists
	existingUser, _ := s.repo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := password.Hash(pwd)
	if err != nil {
		return nil, err
	}

	// Create user entity
	newUser, err := user.NewUser(username, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Login authenticates a user and returns a JWT token
func (s *Service) Login(ctx context.Context, email, pwd string) (string, error) {
	// Find user by email
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Verify password
	if !password.Verify(pwd, u.PasswordHash) {
		return "", ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateToken(u.ID, u.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
