package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID uuid.UUID) (*db.UserEntity, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) AdminAssignRole(ctx context.Context, userID uuid.UUID, role, fullName, subrole string, enterpriseID *uuid.UUID) error {
	validRoles := map[string]bool{"student": true, "brsm": true, "customer": true, "university": true}
	if !validRoles[role] {
		return errors.New("invalid role")
	}

	return s.repo.AssignProfile(ctx, userID, role, fullName, subrole, enterpriseID)
}