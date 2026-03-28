package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type EmploymentService struct {
	repo *repositories.EmploymentRepository
}

func NewEmploymentService(repo *repositories.EmploymentRepository) *EmploymentService {
	return &EmploymentService{repo: repo}
}

func (s *EmploymentService) GetAll(ctx context.Context, status, universityDeptID string, limit, offset int) (*dto.AllEmploymentRequestsResponse, error) {
	requests, total, err := s.repo.GetAll(ctx, status, universityDeptID, limit, offset)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentRequestResponse, len(requests))
	for i, req := range requests {
		response[i] = s.mapToResponse(req)
	}

	return &dto.AllEmploymentRequestsResponse{
		Requests: response,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

func (s *EmploymentService) GetByID(ctx context.Context, id uuid.UUID) (*dto.EmploymentRequestDetailResponse, error) {
	request, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if request == nil {
		return nil, nil
	}

	participantIDs := make([]string, len(request.ParticipantUserIDs))
	for i, p := range request.ParticipantUserIDs {
		participantIDs[i] = p.String()
	}

	return &dto.EmploymentRequestDetailResponse{
		EmploymentRequestResponse: s.mapToResponse(request.EmploymentRequestWithDetails),
		ParticipantUserIDs:        participantIDs,
	}, nil
}

func (s *EmploymentService) Create(ctx context.Context, enterpriseID, universityDeptID uuid.UUID, req *dto.CreateEmploymentRequest) (*dto.EmploymentRequestResponse, error) {
	request, err := s.repo.Create(ctx, enterpriseID, universityDeptID, req.Title, req.Description, req.Requirements, req.Salary, req.Schedule, req.MaxParticipants)
	if err != nil {
		return nil, err
	}

	details, err := s.repo.GetByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	resp := s.mapToResponse(details.EmploymentRequestWithDetails)
	return &resp, nil
}

func (s *EmploymentService) Update(ctx context.Context, id uuid.UUID, req *dto.UpdateEmploymentRequest) error {
	return s.repo.Update(ctx, id, req.Title, req.Description, req.Requirements, req.Salary, req.Schedule, req.MaxParticipants, req.Status)
}

func (s *EmploymentService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *EmploymentService) Apply(ctx context.Context, requestID, userID uuid.UUID) error {
	return s.repo.Apply(ctx, requestID, userID)
}

func (s *EmploymentService) CancelApplication(ctx context.Context, requestID, userID uuid.UUID) error {
	return s.repo.CancelApplication(ctx, requestID, userID)
}

func (s *EmploymentService) UpdateParticipantStatus(ctx context.Context, requestID, userID uuid.UUID, status string) error {
	return s.repo.UpdateParticipantStatus(ctx, requestID, userID, status)
}

func (s *EmploymentService) GetParticipants(ctx context.Context, requestID uuid.UUID) ([]dto.EmploymentParticipantResponse, error) {
	participants, err := s.repo.GetParticipants(ctx, requestID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentParticipantResponse, len(participants))
	for i, p := range participants {
		var contractedAt *string
		if p.ContractedAt != nil {
			t := p.ContractedAt.Format("2006-01-02T15:04:05Z")
			contractedAt = &t
		}

		response[i] = dto.EmploymentParticipantResponse{
			ID:           p.ID.String(),
			RequestID:    p.RequestID.String(),
			UserID:       p.UserID.String(),
			StudentName:  p.StudentName,
			Status:       p.StatusName,
			StatusID:     p.StatusID,
			AppliedAt:    p.AppliedAt.Format("2006-01-02T15:04:05Z"),
			ContractedAt: contractedAt,
		}
	}
	return response, nil
}

func (s *EmploymentService) GetStatuses(ctx context.Context) ([]dto.EmploymentStatusResponse, error) {
	statuses, err := s.repo.GetStatuses(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentStatusResponse, len(statuses))
	for i, st := range statuses {
		response[i] = dto.EmploymentStatusResponse{
			ID:          st.ID,
			Name:        st.Name,
			Description: st.Description,
		}
	}
	return response, nil
}

func (s *EmploymentService) GetParticipantStatuses(ctx context.Context) ([]dto.EmploymentParticipantStatusResponse, error) {
	statuses, err := s.repo.GetParticipantStatuses(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentParticipantStatusResponse, len(statuses))
	for i, st := range statuses {
		response[i] = dto.EmploymentParticipantStatusResponse{
			ID:          st.ID,
			Name:        st.Name,
			Description: st.Description,
		}
	}
	return response, nil
}

func (s *EmploymentService) GetByEnterprise(ctx context.Context, enterpriseID uuid.UUID) ([]dto.EmploymentRequestResponse, error) {
	requests, err := s.repo.GetByEnterprise(ctx, enterpriseID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentRequestResponse, len(requests))
	for i, req := range requests {
		response[i] = s.mapToResponse(req)
	}
	return response, nil
}

func (s *EmploymentService) GetByUniversityDepartment(ctx context.Context, universityDeptID uuid.UUID) ([]dto.EmploymentRequestResponse, error) {
	requests, err := s.repo.GetByUniversityDepartment(ctx, universityDeptID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentRequestResponse, len(requests))
	for i, req := range requests {
		response[i] = s.mapToResponse(req)
	}
	return response, nil
}

func (s *EmploymentService) GetUserApplications(ctx context.Context, userID uuid.UUID) ([]dto.EmploymentApplicationResponse, error) {
	fmt.Println("GetUserApplications called with userID:", userID)
	applications, err := s.repo.GetUserApplications(ctx, userID)
	fmt.Println("GetUserApplications found requests:", len(applications))
	if err != nil {
		return nil, err
	}

	response := make([]dto.EmploymentApplicationResponse, len(applications))
	for i, app := range applications {
		response[i] = dto.EmploymentApplicationResponse{
			EmploymentRequestResponse: s.mapToResponse(app.EmploymentRequestWithDetails),
			ParticipantStatus:         app.ParticipantStatusName,
			ParticipantStatusID:       app.ParticipantStatusID,
			ParticipantAppliedAt:      app.ParticipantAppliedAt.Format("2006-01-02T15:04:05Z"),
		}
	}
	return response, nil
}

func (s *EmploymentService) IsUserApplied(ctx context.Context, requestID, userID uuid.UUID) (bool, error) {
	return s.repo.IsUserApplied(ctx, requestID, userID)
}

func (s *EmploymentService) Approve(ctx context.Context, id uuid.UUID, approved bool) error {
	return s.repo.UpdateStatus(ctx, id, approved)
}

func (s *EmploymentService) mapToResponse(req db.EmploymentRequestWithDetails) dto.EmploymentRequestResponse {
	return dto.EmploymentRequestResponse{
		ID:                     req.ID.String(),
		EnterpriseID:           req.EnterpriseID.String(),
		EnterpriseName:         req.EnterpriseName,
		EnterpriseAddress:      req.EnterpriseAddress,
		UniversityDepartmentID: req.UniversityDepartmentID.String(),
		UniversityName:         req.UniversityName,
		DepartmentName:         req.DepartmentName,
		UniversityAddress:      req.UniversityAddress,
		Title:                  req.Title,
		Description:            req.Description,
		Requirements:           req.Requirements,
		Salary:                 req.Salary,
		Schedule:               req.Schedule,
		MaxParticipants:        req.MaxParticipants,
		CurrentParticipants:    req.CurrentParticipants,
		Status:                 req.StatusName,
		StatusID:               req.StatusID,
		CreatedAt:              req.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:              req.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
