# Nut - Task Management API

Nut is a RESTful API service built in Go that provides ticket/task management capabilities. The service allows users to create, retrieve, update, and list tickets with various attributes like priority, status, and descriptions.

## Current Implementation

### API Endpoints

The following endpoints are currently implemented:

- **Create Ticket**
  - `POST /tickets`
  - Creates a new ticket with title, description, and priority
  - Returns the newly created ticket with a unique ID

- **Get Ticket**
  - `GET /tickets/{ticketId}`
  - Retrieves a specific ticket by ID
  - Returns 404 if ticket not found

- **Update Ticket**
  - `POST /tickets/{ticketId}`
  - Updates an existing ticket's properties
  - Returns the updated ticket information

- **List Tickets**
  - `GET /tickets`
  - Returns all tickets in the system

### Data Model

Tickets have the following properties:
- ID (auto-generated)
- Title
- Description
- Status (open/closed)
- Priority (p1/p2/p3/p4)
- Created At
- Updated At
- Archived flag

### Technical Features

- **Database:** PostgreSQL with pgx driver
- **Configuration:** Environment variable based configuration
- **Architecture:** Clean architecture with separation of:
  - Handlers (HTTP layer)
  - DTOs (Data Transfer Objects)
  - Entities (Domain models)
  - Stores (Data access layer)

### Error Handling
- Proper HTTP status codes (400, 404, 500)
- Structured error responses
- Database error handling

## Setup and Running

### Prerequisites
- Go 1.23.3+
- PostgreSQL 16
- Docker (for local development)

### Environment Variables
```
DB_USER - Database username
DB_PASS - Database password
DB_HOST - Database host
DB_PORT - Database port
DB_NAME - Database name
```

### Local Development

1. Start the database:
```bash
make setup_local
```

2. Run database migrations:
```bash
make migrate_db
```

3. Run the application:
```bash
go run cmd/api/main.go
```

## Roadmap

[] Implementation of `DELETE /tickets/{id}` endpoint
[] Concurrent task processing using goroutines
[] Background task status updates
[] Unit tests for business logic
[] Authentication and Authorization

## Project Structure

```
.
├── cmd/api/           - Application entrypoint
├── docker/            - Docker compose files
├── internal/          - Internal packages
│   ├── config/        - Application configuration
│   ├── constants/     - Shared constants
│   ├── database/      - Database connection
│   ├── dtos/          - Data Transfer Objects
│   ├── entities/      - Domain entities
│   ├── handlers/      - HTTP handlers
│   ├── helpers/       - Utility functions
│   └── stores/        - Data access layer
├── migrations/        - Database migrations
└── Makefile          - Build and development commands
```
