package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/config"
	"github.com/DiaaSaada/crypto-payment-gateway/internal/handler"
	"github.com/DiaaSaada/crypto-payment-gateway/internal/middleware"
	"github.com/DiaaSaada/crypto-payment-gateway/internal/repository/user"
	userUseCase "github.com/DiaaSaada/crypto-payment-gateway/internal/usecase/user"
	"github.com/DiaaSaada/crypto-payment-gateway/pkg/jwt"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize JWT service
	jwtService := jwt.NewService(cfg.JWTSecret, cfg.JWTTokenDuration)

	// Initialize repository (using in-memory for now)
	userRepo := user.NewInMemoryRepository()

	// Initialize use case/service
	userService := userUseCase.NewService(userRepo, jwtService)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Initialize middleware
	authMiddleware := middleware.NewAuth(jwtService)

	// Setup routes
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/api/register", userHandler.Register)
	mux.HandleFunc("/api/login", userHandler.Login)

	// Protected route example
	mux.HandleFunc("/api/protected", authMiddleware.Authenticate(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middleware.UserIDKey)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "You are authenticated", "user_id": "%v"}`, userID)
	}))

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Printf("Available endpoints:")
	log.Printf("  POST /api/register - Register a new user")
	log.Printf("  POST /api/login - Login and get JWT token")
	log.Printf("  GET  /api/protected - Protected endpoint (requires JWT)")
	log.Printf("  GET  /health - Health check")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
