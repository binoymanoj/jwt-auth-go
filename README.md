## JWT-Auth-Go

A simple JWT authentication service built with Go, Gin, and PostgreSQL.

## Overview
JWT-Auth-Go is a lightweight authentication service that provides user registration and authentication using JSON Web Tokens (JWT). The project is built with modern Go practices and leverages the following technologies:

- Gin - High-performance HTTP web framework
- GORM - ORM library for Go
- PostgreSQL - Relational database
- bcrypt - Password hashing
- godotenv - Environment variable management
- air - Live reload for local development

## Features

- User registration (signup)
- User authentication (login) with secure password handling
- JWT token generation and validation
- PostgreSQL database integration with GORM

## Project Structure

```
├── controllers/
│   └── authController.go   # Handles authentication logic
├── initializers/
│   ├── connectToDB.go      # Database connection setup
│   ├── loadEnvVariables.go # Environment configuration
│   └── syncDatabase.go     # Database migration
├── middleware/
│   └── requireAuth.go      # Middleware for validating user auth
├── models/
│   └── userModel.go        # User data model
├── docker-compose.yml      # Docker compose file for PostgreSQL db in local
├── go.mod                  # GO Mod has the dependencies for the project 
├── go.sum                  # Auto generated file while installing required packages
├── main.go                 # Application entry point
├── .env                    # Environment variables (create this file)
├── .air.toml               # Air configuration for hot reload
├── Makefile                # Build automation
└── README.md               # This file
```

## Prerequisites

- [Go](https://golang.org/dl/) 1.19 or higher (using 1.25 in this project - 1.25 recommended)
- [Docker](https://www.docker.com/get-started) and Docker Compose
- [Make](https://www.gnu.org/software/make/) (usually pre-installed on macOS/Linux)

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/jwt-auth-go.git
cd jwt-auth-go
```

## Development Workflow

1. **Start development**
   ```bash
   make db-up    # Start database
   make run      # Run the API
   ```

2. Create a `.env` file in the project root:
   ```bash
   DB_STRING=postgresql://username:password@localhost:5432/jwt_auth_db
   PORT=4000
   JWT_SECRET=your_secret_key_here
   ```

3. **Test your changes**
   ```bash
   make run      # Automatically formats and vets before running
   (or)
   air           # for hot reload, instantly see the updates without building binary each time
   ```

4. **Build for production**
   ```bash
   make build    # Creates binaries in ./bin/ directory
   ```

5. **Clean up**
   ```bash
   make db-down  # Stop database when done
   make clean    # Remove build artifacts
   ```


*air package should be installed already*

## Available Make Commands

Use `make help` to see all available commands:

| Command    | Description                              |
|------------|------------------------------------------|
| `help`     | Show help menu                          |
| `all`      | Format, vet and build the application   |
| `db-up`    | Start the database (Docker Compose)    |
| `db-down`  | Stop the database                       |
| `fmt`      | Format the Go code                      |
| `vet`      | Check for errors using go vet          |
| `run`      | Run the application locally             |
| `build`    | Build binaries for Linux, macOS, Windows |
| `clean`    | Remove temporary and binary directories  |


## How to Use

### Signup

```bash
curl -X POST http://localhost:4000/api/authsignup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"securepassword"}'
```

### Login

```bash
curl -X POST http://localhost:4000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"securepassword"}'
```

### Validate

```bash
curl -X GET http://localhost:4000/api/auth/validate \
  --cookie Authorization=<paste_the_cookie_generated_while_logging_in>
```

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
