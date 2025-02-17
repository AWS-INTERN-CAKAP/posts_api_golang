# POST API GOLANG

# 📌 Introduction

This project is built using Golang and containerized with Docker to ensure a seamless development and deployment experience. It includes user authentication, JWT-based security, PostgreSQL as the database, and follows a clean architecture pattern.

# 📦 Prerequisites

Before running the project, make sure you have the following installed:

- [Docker](https://www.docker.com/)

- [Golang](https://go.dev/doc/install)

- [Migrate](https://github.com/golang-migrate/migrate) (Optional)

# 🛠️ Project Structure

```bash
├── build
│   ├── Dockerfile     # Dockerfile for app container
├── cmd                # Application entry point
├── config             # Configuration files
├── database/migration # Migration files
├── internal           # Core business logic
│   ├── http           # HTTP handlers
│   ├── entity         # Database models
│   ├── repositories   # Database queries
│   ├── services       # Business logic layer
├── pkg                # Utility packages
├── .env               # Environment variables
├── docker-compose.yml # Docker configuration
```

# 🚀 Getting Started

1️⃣ Clone the repository

```bash
git clone https://github.com/AWS-INTERN-CAKAP/posts_api_golang.git
cd post_api_golang
```

2️⃣ Environment Configuration

Create a .env file in the root directory and add the following:

```bash
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
```

3️⃣ Running with Docker

To start the application using Docker Compose, run:

```bash
docker-compose up --build -d
```
This will start the application and the PostgreSQL database.

4️⃣ Running Locally (Without Docker)

If you want to run the application without Docker, follow these steps:

```bash
go mod tidy
migrate -path database/migration/ -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up
go run ./cmd/main.go
```





