package services

import (
	"context"
	"errors"

	"github.com/lanxre/frameby/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo     *repositories.UserRepository
	tokenSvc *TokenService
}

func NewAuthService(repo *repositories.UserRepository, tokenSvc *TokenService) *AuthService {
	return &AuthService{
		repo:     repo,
		tokenSvc: tokenSvc,
	}
}

func (s *AuthService) Register(ctx context.Context, email, login, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = s.repo.Create(ctx, email, login, string(hash))
	return err
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return "", "", errors.New("invalid login or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("invalid login or password")
	}

	token, err := s.tokenSvc.GenerateAccessToken(user.ID, user.Role, user.Email)
	if err != nil {
		return "", "", err
	}

	return token, user.Role, nil
}