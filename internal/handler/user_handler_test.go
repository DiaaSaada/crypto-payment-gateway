package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/handler"
	"github.com/DiaaSaada/crypto-payment-gateway/internal/repository/user"
	userUseCase "github.com/DiaaSaada/crypto-payment-gateway/internal/usecase/user"
	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
)

func setupHandler() *handler.UserHandler {
	repo := user.NewInMemoryRepository()
	jwtService := jwt.NewService("test-secret", 24*time.Hour)
	service := userUseCase.NewService(repo, jwtService)
	return handler.NewUserHandler(service)
}

func TestUserHandler_Register_Success(t *testing.T) {
	h := setupHandler()

	reqBody := handler.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Register(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Register() status = %v, want %v", w.Code, http.StatusCreated)
	}

	var resp handler.RegisterResponse
	json.NewDecoder(w.Body).Decode(&resp)

	if resp.Username != "testuser" {
		t.Errorf("Register() username = %v, want testuser", resp.Username)
	}

	if resp.Email != "test@example.com" {
		t.Errorf("Register() email = %v, want test@example.com", resp.Email)
	}

	if resp.ID == "" {
		t.Error("Register() should return user ID")
	}
}

func TestUserHandler_Register_MissingFields(t *testing.T) {
	h := setupHandler()

	reqBody := handler.RegisterRequest{
		Username: "testuser",
		Email:    "",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Register(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Register() status = %v, want %v", w.Code, http.StatusBadRequest)
	}
}

func TestUserHandler_Register_InvalidMethod(t *testing.T) {
	h := setupHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/register", nil)
	w := httptest.NewRecorder()

	h.Register(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Register() status = %v, want %v", w.Code, http.StatusMethodNotAllowed)
	}
}

func TestUserHandler_Login_Success(t *testing.T) {
	h := setupHandler()

	// First register a user
	regBody := handler.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	regData, _ := json.Marshal(regBody)
	regReq := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(regData))
	regW := httptest.NewRecorder()
	h.Register(regW, regReq)

	// Now login
	loginBody := handler.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	loginData, _ := json.Marshal(loginBody)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(loginData))
	w := httptest.NewRecorder()

	h.Login(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Login() status = %v, want %v", w.Code, http.StatusOK)
	}

	var resp handler.LoginResponse
	json.NewDecoder(w.Body).Decode(&resp)

	if resp.Token == "" {
		t.Error("Login() should return a token")
	}
}

func TestUserHandler_Login_InvalidCredentials(t *testing.T) {
	h := setupHandler()

	loginBody := handler.LoginRequest{
		Email:    "notfound@example.com",
		Password: "password123",
	}
	loginData, _ := json.Marshal(loginBody)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(loginData))
	w := httptest.NewRecorder()

	h.Login(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Login() status = %v, want %v", w.Code, http.StatusUnauthorized)
	}
}

func TestUserHandler_Login_MissingFields(t *testing.T) {
	h := setupHandler()

	loginBody := handler.LoginRequest{
		Email:    "test@example.com",
		Password: "",
	}
	loginData, _ := json.Marshal(loginBody)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(loginData))
	w := httptest.NewRecorder()

	h.Login(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Login() status = %v, want %v", w.Code, http.StatusBadRequest)
	}
}
