package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/dto"
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

func (s *AuthService) Login(ctx context.Context, email, password string) (string, *dto.LoginUserDTO, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	
	if err != nil {
		return "", nil, errors.New("ошибка базы данных")
	}

	if user == nil {
		return "", nil, errors.New("пользователь с таким email не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("неверный пароль")
	}

	token, err := s.tokenSvc.GenerateAccessToken(user.ID, user.Role, user.Email)
	if err != nil {
		return "", nil, err
	}

	userDto := &dto.LoginUserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Login:     user.Login,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return token, userDto, nil
}

func (s *AuthService) GetMe(ctx context.Context, userId uuid.UUID) (*dto.UserDto, error) {
	user, err := s.repo.GetByID(ctx, userId)
	if err != nil {
		fmt.Print("failed to get user by id: %w", err)
		return nil, err
	}

	userDto := &dto.UserDto{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
		Role:  user.Role,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Avatar:    user.Avatar,

		FullName:     &user.FullName,
		Subrole:      &user.Subrole,
		EnterpriseID: user.EnterpriseID,
	}

	return userDto, nil
}
