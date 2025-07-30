# Go Backend Starter API

A clean, modular **Golang REST API** boilerplate with:

- JWT authentication  
- PostgreSQL integration  
- Swagger documentation  
- Layered architecture (Controller → Service → Repository)  
- Unit & integration testing  

---

## Project Structure

```
go-backend/
├── cmd/             # Main application entrypoint
├── config/          # Environment configuration
├── internal/
│   ├── api/         # Controllers, DTOs, Routes
│   ├── auth/        # JWT & middleware
│   ├── common/      # Logging & helpers
│   ├── database/    # DB connection
│   ├── docs/        # Swagger docs (auto-generated)
│   ├── interfaces/  # Interface definitions
│   ├── models/      # Data models
│   ├── repositories/# Data access layer
│   └── services/    # Business logic
├── test/
│   ├── unit/        # Unit tests
│   └── integration/ # Integration tests
├── .env             # Environment variables
├── go.mod / go.sum
└── README.md
```

---

##  Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/MUKUL3115/Go_Backend.git
cd go-backend
```

### 2. Set up environment variables

Create a `.env` file in the root directory:

```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdb
JWT_SECRET=your_jwt_secret_key
```

### 3. Run the server

```bash
go run cmd/main.go
```

The server should start on `http://localhost:8080`

---

## API Documentation (Swagger)

Visit Swagger UI at:  
**`http://localhost:8080/docs/index.html`**

---

## Health Check

Hit `/healthz` to verify the service is running.

---

## Testing

To run unit and integration tests:

```bash
go test ./...
```

---

## Contributions / Hiring

This project was built as a clean base for real-world Golang backends.  
If you're hiring for Go roles — [Let's talk](mailto:mukulraana@gmail.com)!
