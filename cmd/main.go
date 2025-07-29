package main

import (
	"log"
	"os"
	"os/exec"

	"go-backend/config"
	"go-backend/internal/api/routes"
	"go-backend/internal/common"
	"go-backend/internal/database"
	"go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-backend/internal/docs"
)

// @title Go Backend API
// @version 1.0
// @description This is the API documentation for the Go backend service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @host localhost:8082
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token (e.g. "Bearer eyJhbGci...")

func generateSwagger() {
	log.Println("Generating Swagger docs...")
	cmd := exec.Command("swag", "init", "--dir", "./cmd,./internal/api/controllers,./internal/api/dto", "--output", "./internal/docs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to generate Swagger docs: %v", err)
	}
	log.Println("Swagger docs generated.")
}

func main() {
	generateSwagger()

	// Swagger runtime config
	docs.SwaggerInfo.Title = "Go Backend API"
	docs.SwaggerInfo.Description = "This is the API documentation for the Go backend service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8082"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	common.SetupLogger()
	db := database.InitDB()

	gin.SetMode(config.GetEnv("GIN_MODE", "debug")) // or "release"
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r, db)

	r.Run(config.GetEnv("PORT", ":8082"))
}
