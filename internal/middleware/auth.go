package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
)

type contextKey string

const UserIDKey contextKey = "user_id"

// Auth is a middleware that validates JWT tokens
type Auth struct {
	jwtService *jwt.Service
}

// NewAuth creates a new authentication middleware
func NewAuth(jwtService *jwt.Service) *Auth {
	return &Auth{
		jwtService: jwtService,
	}
}

// Authenticate validates the JWT token and adds user info to context
func (a *Auth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			a.sendError(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			a.sendError(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		// Validate token
		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			a.sendError(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (a *Auth) sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
