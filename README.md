# Veediver 

A Go-based REST API server built with the Chi framework.

## Features

- RESTful API endpoints
- Health check endpoint
- Structured project layout
- JSON responses
- CRUD operations (placeholder implementation)

## Project Structure

```
seek-clarity/
├── main.go           # Main application entry point
├── handlers/         # HTTP request handlers
├── middleware/       # Custom middleware
├── models/          # Data models
├── routes/          # Route definitions
├── config/          # Configuration files
├── go.mod           # Go module file
└── README.md        # This file
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git

### Installation

1. Navigate to the project directory:
   ```bash
   cd seek-clarity
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Returns server health status

### API v1
- `GET /api/v1/` - Welcome message
- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user by ID
- `DELETE /api/v1/users/{id}` - Delete user by ID

## Development

### Building the application

```bash
go build -o seek-clarity main.go
```

### Running tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License.
