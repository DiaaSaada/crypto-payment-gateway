package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/repository/user"
	userUseCase "github.com/DiaaSaada/crypto-payment-gateway/internal/usecase/user"
	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
)

func TestService_Register(t *testing.T) {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)

	ctx := context.Background()

	u, err := service.Register(ctx, "testuser", "test@example.com", "password123")
	if err != nil {
		t.Errorf("Register() unexpected error = %v", err)
	}

	if u.Username != "testuser" {
		t.Errorf("Register() username = %v, want testuser", u.Username)
	}

	if u.Email != "test@example.com" {
		t.Errorf("Register() email = %v, want test@example.com", u.Email)
	}

	if u.ID == "" {
		t.Error("Register() should set user ID")
	}
}

func TestService_Register_DuplicateEmail(t *testing.T) {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)

	ctx := context.Background()

	_, _ = service.Register(ctx, "testuser1", "test@example.com", "password123")
	_, err := service.Register(ctx, "testuser2", "test@example.com", "password123")

	if err == nil {
		t.Error("Register() expected error for duplicate email")
	}
}

func TestService_Login_Success(t *testing.T) {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)

	ctx := context.Background()

	// Register a user first
	_, _ = service.Register(ctx, "testuser", "test@example.com", "password123")

	// Try to login
	token, err := service.Login(ctx, "test@example.com", "password123")
	if err != nil {
		t.Errorf("Login() unexpected error = %v", err)
	}

	if token == "" {
		t.Error("Login() should return a token")
	}
}

func TestService_Login_InvalidEmail(t *testing.T) {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)

	ctx := context.Background()

	_, err := service.Login(ctx, "notfound@example.com", "password123")

	if err != userUseCase.ErrInvalidCredentials {
		t.Errorf("Login() expected ErrInvalidCredentials, got %v", err)
	}
}

func TestService_Login_InvalidPassword(t *testing.T) {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)

	ctx := context.Background()

	// Register a user first
	_, _ = service.Register(ctx, "testuser", "test@example.com", "password123")

	// Try to login with wrong password
	_, err := service.Login(ctx, "test@example.com", "wrongpassword")

	if err != userUseCase.ErrInvalidCredentials {
		t.Errorf("Login() expected ErrInvalidCredentials, got %v", err)
	}
}
