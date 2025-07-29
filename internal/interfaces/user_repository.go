package interfaces

import (
	"context"
	"go-backend/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
