package services

import (
	"context"
	"errors"
	"go-backend/internal/interfaces"
	"go-backend/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) interfaces.UserService {
	return &userService{repo: r}
}

func (s *userService) Register(ctx context.Context, user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return s.repo.Create(ctx, user)
}

func (s *userService) Login(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil || u == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return u, nil
}
