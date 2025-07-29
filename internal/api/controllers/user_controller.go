package controllers

import (
	"context"
	"net/http"

	"go-backend/internal/api/dto"
	"go-backend/internal/auth"
	"go-backend/internal/common"
	"go-backend/internal/interfaces"
	"go-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.

// UserController handles user-related operations
type UserController struct {
	Service interfaces.UserService
}

// NewUserController creates a new UserController
func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{Service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "User registration data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.JSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{Email: req.Email, Password: req.Password}
	if err := uc.Service.Register(context.Background(), user); err != nil {
		common.JSON(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	common.JSON(c, http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (uc *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.JSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.Service.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		common.JSON(c, http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := auth.GenerateJWT(user.Email)
	common.JSON(c, http.StatusOK, dto.AuthResponse{Token: token})
}

// Profile godoc
// @Summary Get user profile
// @Description Returns basic profile info for the logged-in user
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /secure/profile [get]
func (uc *UserController) Profile(c *gin.Context) {
	email := c.GetString("userEmail") // Extracted from JWT middleware
	common.JSON(c, http.StatusOK, gin.H{"email": email})
}
