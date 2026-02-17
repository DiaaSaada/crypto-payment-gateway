# Crypto Payment Gateway

A web-based REST API built with Go following Domain-Driven Design (DDD) and Test-Driven Development (TDD) principles.

## Architecture

This project follows a clean architecture with clear separation of concerns:

### Layers

1. **Domain Layer** (`internal/domain/`): Contains business entities and repository interfaces
2. **Use Case Layer** (`internal/usecase/`): Contains business logic and use cases
3. **Repository Layer** (`internal/repository/`): Implements data persistence (currently in-memory)
4. **Handler Layer** (`internal/handler/`): HTTP request handlers
5. **Middleware Layer** (`internal/middleware/`): HTTP middleware (authentication, etc.)
6. **Package Layer** (`pkg/`): Reusable utilities (JWT, password hashing)

### Request Flow

```
HTTP Request → Route → Handler → UseCase/Service → Repository → Domain Entity
```

## Features

- ✅ User Registration with password hashing
- ✅ JWT-based Authentication
- ✅ Protected endpoints with JWT middleware
- ✅ Clean architecture following DDD principles
- ✅ Comprehensive unit and integration tests
- ✅ In-memory data storage

## Project Structure

```
crypto-payment-gateway/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── domain/
│   │   └── user/
│   │       ├── user.go            # User domain entity
│   │       ├── user_test.go       # Domain tests
│   │       └── repository.go      # Repository interface (abstract)
│   ├── usecase/
│   │   └── user/
│   │       ├── service.go         # User business logic
│   │       └── service_test.go    # Use case tests
│   ├── repository/
│   │   └── user/
│   │       ├── inmemory.go        # In-memory repository implementation
│   │       └── inmemory_test.go   # Repository tests
│   ├── handler/
│   │   ├── user_handler.go        # HTTP handlers
│   │   └── user_handler_test.go   # Handler integration tests
│   └── middleware/
│       └── auth.go                # JWT authentication middleware
├── pkg/
│   ├── jwt/
│   │   ├── jwt.go                 # JWT token generation/validation
│   │   └── jwt_test.go            # JWT tests
│   └── password/
│       ├── password.go            # Password hashing utilities
│       └── password_test.go       # Password tests
├── go.mod                         # Go module dependencies
└── README.md                      # This file
```

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/DiaaSaada/crypto-payment-gateway.git
cd crypto-payment-gateway
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o bin/api ./cmd/api
```

### Running the Application

```bash
./bin/api
```

Or directly with Go:
```bash
go run cmd/api/main.go
```

The server will start on port 8080 by default.

### Configuration

Environment variables:
- `PORT`: Server port (default: 8080)
- `JWT_SECRET`: Secret key for JWT signing (default: "your-secret-key-change-in-production")
- `JWT_DURATION`: JWT token duration in hours (default: 24)

## API Endpoints

### Health Check

```bash
GET /health
```

Response:
```json
{
  "status": "healthy"
}
```

### User Registration

```bash
POST /api/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

Response:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "john_doe",
  "email": "john@example.com",
  "message": "User registered successfully"
}
```

### User Login

```bash
POST /api/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Protected Endpoint (Example)

```bash
GET /api/protected
Authorization: Bearer <your-jwt-token>
```

Response:
```json
{
  "message": "You are authenticated",
  "user_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test ./... -cover
```

Run tests with verbose output:
```bash
go test ./... -v
```

## Development

### Design Principles

1. **Domain-Driven Design (DDD)**:
   - Domain entities in `internal/domain/`
   - Business logic in use cases/services
   - Repository pattern for data access
   - Clear bounded contexts

2. **Test-Driven Development (TDD)**:
   - Comprehensive unit tests for all layers
   - Integration tests for HTTP handlers
   - High test coverage (>80%)

3. **Modular Architecture**:
   - Clear separation of concerns
   - Dependency injection
   - Interface-based design
   - Easy to extend and maintain

4. **Clean Code**:
   - Single Responsibility Principle
   - Interface Segregation
   - Dependency Inversion

### Adding New Features

To add a new feature:

1. Define domain entities in `internal/domain/`
2. Create repository interface in the domain package
3. Implement repository in `internal/repository/`
4. Create use case/service in `internal/usecase/`
5. Add HTTP handlers in `internal/handler/`
6. Register routes in `cmd/api/main.go`
7. Write tests for all layers

## Dependencies

- [github.com/google/uuid](https://github.com/google/uuid) - UUID generation
- [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) - JWT authentication
- [golang.org/x/crypto](https://golang.org/x/crypto) - Password hashing (bcrypt)

## License

This project is part of a crypto payment gateway system.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for your changes
4. Implement your changes
5. Ensure all tests pass
6. Submit a pull request