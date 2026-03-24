package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type UniversityDepartmentService struct {
	repo *repositories.UniversityDepartmentRepository
}

func NewUniversityDepartmentService(repo *repositories.UniversityDepartmentRepository) *UniversityDepartmentService {
	return &UniversityDepartmentService{repo: repo}
}

func (s *UniversityDepartmentService) GetAll(ctx context.Context) ([]dto.UniversityDepartmentResponse, error) {
	departments, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.UniversityDepartmentResponse, len(departments))
	for i, d := range departments {
		response[i] = dto.UniversityDepartmentResponse{
			ID:             d.ID.String(),
			UniversityName: d.UniversityName,
			DepartmentName: d.DepartmentName,
			Address:        d.Address,
		}
	}
	return response, nil
}

func (s *UniversityDepartmentService) GetByID(ctx context.Context, id uuid.UUID) (*dto.UniversityDepartmentResponse, error) {
	department, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if department == nil || department.University == nil || department.Department == nil {
		return nil, nil
	}
	return &dto.UniversityDepartmentResponse{
		ID:             department.ID.String(),
		UniversityName: department.University.Name,
		DepartmentName: department.Department.Name,
		Address:        department.Address,
	}, nil
}

func (s *UniversityDepartmentService) Create(ctx context.Context, universityName, departmentName string, address *string) (*dto.UniversityDepartmentResponse, error) {
	department, err := s.repo.Create(ctx, universityName, departmentName, address)
	if err != nil {
		return nil, err
	}
	return &dto.UniversityDepartmentResponse{
		ID:             department.ID.String(),
		UniversityName: department.UniversityName,
		DepartmentName: department.DepartmentName,
		Address:        department.Address,
	}, nil
}
