# 🧠 Go Backend Starter API

A clean, modular **Golang REST API** boilerplate with:

- ✅ JWT authentication  
- ✅ PostgreSQL integration  
- ✅ Swagger documentation  
- ✅ Layered architecture (Controller → Service → Repository)  
- ✅ Unit & integration testing  

---

## 📁 Project Structure

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

## 🚀 Setup Instructions

### 1. Clone & Install Dependencies

```bash
git clone https://github.com/yourusername/go-backend.git
cd go-backend
go mod tidy
```

### 2. Create `.env`

Example:

```
PORT=:8081
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
JWT_SECRET=your_secret_key
GIN_MODE=debug
```

---

## 🚀 Running the Server

```bash
go run cmd/main.go
```
Or build and run:

```bash
go build -o server ./cmd
./server
```

---

## 🧪 Running Tests

### Unit Tests

```bash
go test ./test/unit/...
```

### Integration Tests

> Ensure the DB is running:

```bash
go test ./test/integration/...
```

---

## 📚 Swagger API Docs

### Generate Swagger Docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init --dir ./cmd,./internal/api --output ./internal/docs
```

### Access Swagger UI

Open in browser:

```
http://localhost:8081/swagger/index.html
```

---

## ✅ Sample API Usage

### Register

```
POST /api/v1/register
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "yourpassword"
}
```

### Login

```
POST /api/v1/login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "yourpassword"
}
```

### Secure Profile (requires JWT)

```
GET /api/v1/secure/profile
Authorization: Bearer <your_token>
```

---

## 🧰 Tools & Libraries

- Gin
- GORM
- Swaggo
- PostgreSQL
- bcrypt
- godotenv

---

## 👨‍💻 Author

**Mukul Rana**

Feel free to fork, contribute or raise issues!