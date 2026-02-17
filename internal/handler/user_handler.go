package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/usecase/user"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userUseCase user.UseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase user.UseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse represents the registration response
type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Register handles user registration
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Username == "" || req.Email == "" || req.Password == "" {
		h.sendError(w, "Username, email and password are required", http.StatusBadRequest)
		return
	}

	// Call use case
	newUser, err := h.userUseCase.Register(r.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send response
	resp := RegisterResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Message:  "User registered successfully",
	}
	h.sendJSON(w, resp, http.StatusCreated)
}

// Login handles user authentication
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Email == "" || req.Password == "" {
		h.sendError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Call use case
	token, err := h.userUseCase.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send response
	resp := LoginResponse{
		Token: token,
	}
	h.sendJSON(w, resp, http.StatusOK)
}

// Helper methods
func (h *UserHandler) sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) sendError(w http.ResponseWriter, message string, status int) {
	h.sendJSON(w, ErrorResponse{Error: message}, status)
}
