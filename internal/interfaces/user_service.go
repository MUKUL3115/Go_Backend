package interfaces

import (
	"context"
	"go-backend/internal/models"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (*models.User, error)
}
