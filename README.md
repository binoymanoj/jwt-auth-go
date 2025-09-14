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
└── README.md               # This file
```

## Prerequisites

- Go 1.16+
- Docker
- Make
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/jwt-auth-go.git
cd jwt-auth-go
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the project root:

```bash
DB_STRING=postgresql://username:password@localhost:5432/jwt_auth_db
PORT=4000
JWT_SECRET=your_secret_key_here
```

4. Create/Run the PostgreSQL database (using docker):

```bash
make db-up
(or)
docker compose up -d
```

5. Build and run the application:

a. Standard mode:

```bash
make build
(or)
make
```

b. With hot reload (recommended for development)

```bash
air
```

*air package should be installed already*

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
