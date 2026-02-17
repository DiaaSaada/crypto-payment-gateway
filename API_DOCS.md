# API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Health Check

Check if the API is running.

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

---

### 2. User Registration

Register a new user account.

**Endpoint:** `POST /api/register`

**Request:**
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

**Response (Success - 201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "john_doe",
  "email": "john@example.com",
  "message": "User registered successfully"
}
```

**Response (Error - 400):**
```json
{
  "error": "Username, email and password are required"
}
```

```json
{
  "error": "user with this email already exists"
}
```

---

### 3. User Login

Authenticate and get a JWT token.

**Endpoint:** `POST /api/login`

**Request:**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

**Response (Success - 200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (Error - 401):**
```json
{
  "error": "invalid credentials"
}
```

---

### 4. Protected Endpoint (Example)

Access a protected endpoint using JWT token.

**Endpoint:** `GET /api/protected`

**Headers:**
- `Authorization: Bearer <jwt-token>`

**Request:**
```bash
# First, get the token
TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword123"
  }' | jq -r '.token')

# Then use it to access protected endpoint
curl -X GET http://localhost:8080/api/protected \
  -H "Authorization: Bearer $TOKEN"
```

**Response (Success - 200):**
```json
{
  "message": "You are authenticated",
  "user_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Response (Error - 401):**
```json
{
  "error": "Missing authorization header"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

---

## Complete Example Workflow

### 1. Register a new user
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "mypassword123"
  }'
```

### 2. Login and save token
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "mypassword123"
  }' | jq -r '.token')

echo "Token: $TOKEN"
```

### 3. Access protected resources
```bash
curl -X GET http://localhost:8080/api/protected \
  -H "Authorization: Bearer $TOKEN"
```

---

## Error Responses

All error responses follow this format:
```json
{
  "error": "error message here"
}
```

### Common HTTP Status Codes

- `200 OK` - Request succeeded
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Authentication failed
- `405 Method Not Allowed` - Wrong HTTP method
- `500 Internal Server Error` - Server error

---

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. After logging in, include the token in the `Authorization` header for protected endpoints:

```
Authorization: Bearer <your-jwt-token>
```

Tokens expire after 24 hours by default (configurable via `JWT_DURATION` environment variable).
