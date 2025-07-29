package routes

import (
	"go-backend/internal/api/controllers"
	"go-backend/internal/auth"
	"go-backend/internal/repositories"
	"go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1")

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	v1.POST("/register", userController.Register)
	v1.POST("/login", userController.Login)

	protected := v1.Group("/secure")
	protected.Use(auth.JWTAuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Secure route accessed"})
	})
}
