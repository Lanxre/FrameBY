package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type SpecialtyService struct {
	repo *repositories.SpecialtyRepository
}

func NewSpecialtyService(repo *repositories.SpecialtyRepository) *SpecialtyService {
	return &SpecialtyService{repo: repo}
}

func mapSpecialty(e db.SpecialtyEntity) dto.SpecialtyResponse {
	return dto.SpecialtyResponse{
		ID:           e.ID.String(),
		Name:         e.Name,
		DepartmentID: e.DepartmentID.String(),
	}
}

func mapSpecialties(entities []db.SpecialtyEntity) []dto.SpecialtyResponse {
	result := make([]dto.SpecialtyResponse, len(entities))
	for i, e := range entities {
		result[i] = mapSpecialty(e)
	}
	return result
}

func (s *SpecialtyService) Create(ctx context.Context, name string, departmentID uuid.UUID) (*dto.SpecialtyResponse, error) {
	entity, err := s.repo.Create(ctx, name, departmentID)
	if err != nil {
		return nil, err
	}
	result := mapSpecialty(*entity)
	return &result, nil
}

func (s *SpecialtyService) GetByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]dto.SpecialtyResponse, error) {
	entities, err := s.repo.GetByDepartmentID(ctx, departmentID)
	if err != nil {
		return nil, err
	}
	return mapSpecialties(entities), nil
}

func (s *SpecialtyService) GetByUniversityDepartmentID(ctx context.Context, universityDepartmentID uuid.UUID) ([]dto.SpecialtyResponse, error) {
	entities, err := s.repo.GetByUniversityDepartmentID(ctx, universityDepartmentID)
	if err != nil {
		return nil, err
	}
	return mapSpecialties(entities), nil
}
