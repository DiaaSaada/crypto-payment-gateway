package jwt_test

import (
	"testing"
	"time"

	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
)

func TestGenerateToken(t *testing.T) {
	service := jwt.NewService("test-secret-key", 24*time.Hour)

	token, err := service.GenerateToken("user123", "test@example.com")
	if err != nil {
		t.Errorf("GenerateToken() unexpected error = %v", err)
	}

	if token == "" {
		t.Error("GenerateToken() returned empty token")
	}
}

func TestValidateToken_ValidToken(t *testing.T) {
	service := jwt.NewService("test-secret-key", 24*time.Hour)

	token, _ := service.GenerateToken("user123", "test@example.com")

	claims, err := service.ValidateToken(token)
	if err != nil {
		t.Errorf("ValidateToken() unexpected error = %v", err)
	}

	if claims.UserID != "user123" {
		t.Errorf("ValidateToken() UserID = %v, want user123", claims.UserID)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("ValidateToken() Email = %v, want test@example.com", claims.Email)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	service := jwt.NewService("test-secret-key", 24*time.Hour)

	_, err := service.ValidateToken("invalid-token")
	if err == nil {
		t.Error("ValidateToken() expected error for invalid token")
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	service := jwt.NewService("test-secret-key", -1*time.Hour)

	token, _ := service.GenerateToken("user123", "test@example.com")

	_, err := service.ValidateToken(token)
	if err == nil {
		t.Error("ValidateToken() expected error for expired token")
	}
}
