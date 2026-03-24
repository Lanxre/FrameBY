package services

import (
	"context"

	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type EnterpriseService struct {
	repo *repositories.EnterpriseRepository
}

func NewEnterpriseService(repo *repositories.EnterpriseRepository) *EnterpriseService {
	return &EnterpriseService{repo: repo}
}

func (s *EnterpriseService) GetAll(ctx context.Context) ([]dto.EnterpriseResponse, error) {
	enterprises, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EnterpriseResponse, len(enterprises))
	for i, e := range enterprises {
		response[i] = dto.EnterpriseResponse{
			ID:      e.ID.String(),
			Name:    e.Name,
			Address: e.Address,
		}
	}
	return response, nil
}
