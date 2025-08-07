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

	// Import generated Swagger docs
	_ "go-backend/internal/docs"
	"go-backend/internal/docs"
)

// @title           Go Backend API
// @version         1.0
// @description     This is the API documentation for the Go backend service.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.email   support@example.com

// @host      localhost:8082
// @BasePath  /api/v1
// @schemes   http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token (e.g. "Bearer eyJhbGci...")

func generateSwagger() {
	log.Println("Generating Swagger docs...")
	cmd := exec.Command("swag", "init",
		"--dir", "./cmd,./internal/api/controllers,./internal/api/dto",
		"--output", "./internal/docs",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to generate Swagger docs: %v", err)
	}
	log.Println("Swagger docs generated.")
}

func main() {
	// Step 1: Generate Swagger docs
	generateSwagger()

	// Step 2: Swagger runtime metadata
	docs.SwaggerInfo.Title = "Go Backend API"
	docs.SwaggerInfo.Description = "This is the API documentation for the Go backend service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8082"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Step 3: Setup logging and DB
	common.SetupLogger()
	db := database.InitDB()

	// Step 4: Setup Gin
	gin.SetMode(config.GetEnv("GIN_MODE", "debug"))
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	routes.RegisterRoutes(r, db)

	// Step 5: Run server
	port := config.GetEnv("PORT", ":8082")
	log.Printf("Server running at http://localhost%s\n", port)
	r.Run(port)
}
