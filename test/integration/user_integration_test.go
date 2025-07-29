package integration

import (
	"bytes"
	"encoding/json"
	"go-backend/internal/database"
	"go-backend/internal/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../../.env")
	db := database.InitDB()
	r := gin.Default()
	routes.RegisterRoutes(r, db)
	return r
}

func TestUserRegistrationIntegration(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{
		"email":    "integration@test.com",
		"password": "integration123",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}
}
