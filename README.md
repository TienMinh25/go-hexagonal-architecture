# Go Hexagonal Architecture

A **Simple RESTful Service API** written in Go implementing **Hexagonal Architecture** (Ports and Adapters pattern) using modern Go technologies and best practices.

![Architecture Diagram](./docs/architecture.webp)

## 🏗️ Architecture Overview

This project demonstrates the implementation of **Hexagonal Architecture** in Go, also known as the **Ports and Adapters** pattern. The architecture provides a clear separation of concerns and makes the application highly testable and maintainable.

### Key Principles

- **Domain-Driven Design (DDD)**: Business logic is isolated in the domain layer
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Testability**: Each layer can be tested in isolation
- **Flexibility**: Easy to swap implementations (e.g., database, web framework)

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: [Gin](https://gin-gonic.com/) - Fast HTTP web framework
- **Database**: PostgreSQL with native driver
- **Cache**: Redis for caching and session management
- **Authentication**: [PASETO](https://paseto.io/) - Secure token-based authentication
- **Configuration**: Environment variables with structured config
- **Logging**: Structured logging with [slog](https://pkg.go.dev/log/slog)
- **API Documentation**: Swagger with [gin-swagger](https://github.com/swaggo/gin-swagger)
- **Validation**: Request validation with [go-playground/validator](https://github.com/go-playground/validator)

## 📁 Project Structure

```
go-hexagonal-architecture/
├── cmd/
│   └── http/
│       └── main.go              # Application entry point
├── internal/
│   ├── application/
│   │   ├── domain/              # Domain entities and business rules
│   │   │   ├── auth/            # Domain auth
│   │   │   ├── category/        # Domain category
│   │   │   ├── order/           # Domain order
│   │   │   ├── payment/         # Domain payment
│   │   │   ├── product/         # Domain product
│   │   │   ├── user/            # Domain user
│   │   │   └── error.go         # Domain error
│   │   ├── mock/                # Application mock ports/interfaces
│   │   ├── ports/               # Application ports/interfaces
│   │   │   ├── in/              # Primary/Driving ports
│   │   │   └── out/             # Secondary/Driven ports
│   │   ├── util/                # Application util
│   │   └── usecase/             # Application services (use cases)
│   ├── adapter/
│   │   ├── auth/paseto/         # Auth (primary adapter)
│   │   ├── config/              # Config for application
│   │   ├── handler/http/        # HTTP handlers (primary adapter)
│   │   ├── logger/              # Logger util for application
│   │   └── storage/             # Storage implementations
│   │   │   ├── postgres/        # Secondary adapater for postgres
│   │   │   ├── redis/           # Secondary adapater for redis
├── docs/
│   ├── architecture.webp       # Architecture diagram
│   └── db.webp                 # Database schema
├── docker-compose.yml           # Local development setup
├── Dockerfile                   # Container image
├── .env.example                 # Environment variables template
├── go.mod                       # Go modules
├── go.sum                       # Go modules checksums
├── LICENSE                      
├── Makefile                      
└── README.md                    # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for local development)
- PostgreSQL (if running without Docker)
- Redis (if running without Docker)

### Local Development with Docker

1. **Clone the repository**
   ```bash
   git clone https://github.com/TienMinh25/go-hexagonal-architecture.git
   cd go-hexagonal-architecture
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start services with Docker Compose**
   ```bash
   docker-compose up -d
   ```

4. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

## 📊 Database Schema

![Database Schema](./docs/db.webp)

The database schema includes the core entities and their relationships, designed following DDD principles.

## 🔌 API Endpoints

### Users
```http
GET    /api/v1/users          # Get all users
GET    /api/v1/users/:id      # Get user by ID
POST   /api/v1/users          # Create new user
PUT    /api/v1/users/:id      # Update user
DELETE /api/v1/users/:id      # Delete user
```

### Example Request
```bash
# Create a new user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

## 🧪 Testing

The hexagonal architecture makes testing straightforward by allowing each layer to be tested independently.

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

### Test Structure
- **Unit Tests**: Test individual components in isolation

## 🏛️ Architecture Layers

### 1. Domain Layer (Core)
Contains the business entities and rules. This layer has no dependencies on external frameworks or libraries.

- **Entities**: Core business objects
- **Value Objects**: Immutable objects that represent concepts
- **Domain Services**: Business logic that doesn't belong to a single entity

### 2. Application Layer (Core)
Contains application-specific business logic and use cases.

- **Services**: Application services that orchestrate domain operations
- **Ports**: Interfaces that define contracts with the outside world

### 3. Adapter Layer
Implements the ports defined in the application layer.

- **Primary Adapters**: Handle incoming requests (HTTP handlers)
- **Secondary Adapters**: Handle outgoing requests (database, external APIs)

### 4. Infrastructure Layer
Provides technical capabilities that support the higher layers.

- **Configuration**: Application configuration management
- **Database**: Database connection and setup
- **Logging**: Application logging setup

## 🐳 Docker Support

### Build Docker Image
```bash
docker build -t go-hexagonal-architecture .
```

### Run with Docker Compose
```bash
docker-compose up --build
```

The `docker-compose.yml` includes:
- Application container
- PostgreSQL database
- Redis cache
- pgAdmin (database management)

## ⚙️ Configuration

Environment variables can be set in `.env` file or as system environment variables:

```bash
# Server Configuration
PORT=8080
GIN_MODE=release

# Database Configuration
DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_MAX_IDLE_TIME=15m

# Redis Configuration
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

## 🚀 Deployment

### Production Build
```bash
# Build optimized binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/server/main.go

# Or use Docker
docker build -t go-hexagonal-architecture:latest .
```

### Deployment Options
- **Container Orchestration**: Kubernetes, Docker Swarm
- **Cloud Platforms**: AWS ECS, Google Cloud Run, Azure Container Instances
- **Traditional Servers**: Linux VPS with reverse proxy (nginx)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style
- Follow Go conventions and best practices
- Use `gofmt` for formatting
- Run `golint` and `go vet` before committing
- Write tests for new features

## 📝 License

This project is licensed under the Apache License - see the [LICENSE](https://github.com/TienMinh25/go-hexagonal-architecture/blob/main/LICENSE) file for details.

## 🙏 Acknowledgments

- [Hexagonal Architecture by Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Hexagonal Architecture Deep Dive](https://www.happycoders.eu/software-craftsmanship/hexagonal-architecture/)
- [Go community](https://golang.org/community) for excellent tools and libraries

## 📚 Additional Resources

- [Hexagonal Architecture Pattern](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
- [Go Best Practices](https://golang.org/doc/effective_go.html)
- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)

---

**Built with ❤️ using Go and Hexagonal Architecture principles**