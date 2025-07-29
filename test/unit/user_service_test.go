package unit

import (
	"context"
	"go-backend/internal/models"
	"go-backend/internal/services"
	"testing"
)

type mockRepo struct{}

func (m *mockRepo) Create(ctx context.Context, user *models.User) error {
	return nil
}

func (m *mockRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{
		Email:    email,
		Password: "$2a$10$Qt1Q6i23yi2k1TSb6J/wFO3gXshh06DGEc0LiZnkQtcgGzwNwfY3C", 
	}, nil
}

func TestLogin(t *testing.T) {
	svc := services.NewUserService(&mockRepo{})
	user, err := svc.Login(context.Background(), "integration@test.com", "integration123")
	if err != nil || user.Email != "integration@test.com" {
		t.Fatalf("expected user to be returned, got error: %v", err)
	}
}
