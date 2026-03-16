# Warehouse Backend

REST API project built with Go (Gin Gonic), PostgreSQL, and Clean Architecture.

## Architecture

This project follows **Clean Architecture** principles:
- **Domain**: Entities and repository/usecase interfaces.
- **Usecase**: Business logic implementation.
- **Repository**: Data access implementation (PostgreSQL).
- **Delivery**: HTTP handlers and routing (Gin).

## Module Structure

- `cmd/api`: Entry point for the REST API.
- `internal/`: Private application code.
- `pkg/`: Shared utility packages.
- `migrations/`: SQL migration files.

## Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose

### Run Locally

1. Create `.env` file from `.env.example`:
   ```bash
   cp .env.example .env
   ```

2. Start the database:
   ```bash
   docker-compose up -d postgres
   ```

3. Run the application:
   ```bash
   make run
   ```

## API Endpoints

### Health Check
- `GET /health`

### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
