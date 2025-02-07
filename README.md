POST API GOLANG

📌 Introduction

This project is built using Golang and containerized with Docker to ensure a seamless development and deployment experience. It includes user authentication, JWT-based security, PostgreSQL as the database, and follows a clean architecture pattern.

📦 Prerequisites

Before running the project, make sure you have the following installed:

Docker

Docker Compose

Golang

🚀 Getting Started

1️⃣ Clone the repository

git clone [ https://github.com/your-repo.git](https://github.com/AWS-INTERN-CAKAP/posts_api_golang.git)
cd post_api_golang

2️⃣ Environment Configuration

Create a .env file in the root directory and add the following:

# Environment
ENV=dev

# Server Configuration
PORT=8080

# PostgreSQL Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DATABASE=postgres

# JWT Secret Key
JWT_SECRET_KEY=your-secret-key

3️⃣ Running with Docker

To start the application using Docker Compose, run:

docker-compose up --build -d

This will start the application and the PostgreSQL database.

4️⃣ Running Locally (Without Docker)

If you want to run the application without Docker, follow these steps:

go mod tidy
go run ./cmd/main.go

🛠️ Project Structure

├── build
│   ├── Dockerfile    # Dockerfile for app container
├── cmd               # Application entry point
├── config            # Configuration files
├── internal          # Core business logic
│   ├── http          # HTTP handlers
│   ├── middleware    # Middleware logic
│   ├── entity        # Database models
│   ├── repository    # Database queries
│   ├── services      # Business logic layer
├── pkg               # Utility packages
├── .env              # Environment variables
├── docker-compose.yml # Docker configuration



