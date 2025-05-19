## JWT-Auth-Go

A simple JWT authentication service built with Go, Gin, and PostgreSQL.

## Overview
JWT-Auth-Go is a lightweight authentication service that provides user registration and authentication using JSON Web Tokens (JWT). The project is built with modern Go practices and leverages the following technologies:

- Gin - High-performance HTTP web framework
- GORM - ORM library for Go
- PostgreSQL - Relational database
- bcrypt - Password hashing
- godotenv - Environment variable management

## Features

- User registration (signup)
- User authentication (login) with secure password handling
- JWT token generation and validation
- PostgreSQL database integration with GORM

## Project Structure

├── controllers/
│   └── authController.go   # Handles authentication logic
├── initializers/
│   ├── connectToDB.go      # Database connection setup
│   ├── loadEnvVariables.go # Environment configuration
│   └── syncDatabase.go     # Database migration
├── models/
│   └── userModel.go        # User data model
├── main.go                 # Application entry point
├── .env                    # Environment variables (create this file)
└── README.md               # This file

## Prerequisites

- Go 1.16+
- PostgreSQL
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

4. Create the PostgreSQL database:

```bash
createdb jwt_auth_db
```

5. Build and run the application:

```bash
go run main.go
```

## How to Use

### Signup

```bash
curl -X POST http://localhost:4000/signup \
  -H "Content-Type: application/json" \
  -d '{"Email":"user@example.com","Password":"securepassword"}'
```

### Login (Once implemented)

```bash
bashcurl -X POST http://localhost:4000/login \
  -H "Content-Type: application/json" \
  -d '{"Email":"user@example.com","Password":"securepassword"}'
```

## Development
### To Do

- Complete the login endpoint implementation to generate JWT tokens
- Add middleware for JWT token validation
- Create protected routes that require authentication
- Add user profile management
- Implement token refresh functionality

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
