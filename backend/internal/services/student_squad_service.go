package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type StudentSquadService struct {
	repo *repositories.StudentSquadRepository
}

func NewStudentSquadService(repo *repositories.StudentSquadRepository) *StudentSquadService {
	return &StudentSquadService{repo: repo}
}

func mapParticipants(dbParticipants []db.SquadParticipantInfo) []dto.SquadParticipantInfo {
	if len(dbParticipants) == 0 {
		return []dto.SquadParticipantInfo{}
	}
	res := make([]dto.SquadParticipantInfo, len(dbParticipants))
	for i, p := range dbParticipants {
		res[i] = dto.SquadParticipantInfo{
			ID:        p.UserID.String(),
			FullName:  p.FullName,
			Phone:     p.Phone,
			Specialty: p.Specialty,
			Grade:     p.Grade,
			Avatar:    p.Avatar,
		}
	}
	return res
}

func (s *StudentSquadService) GetAll(ctx context.Context, status string, limit, offset int, universityDeptID *uuid.UUID) (*dto.AllSquadsResponse, error) {
	squads, total, err := s.repo.GetAll(ctx, status, limit, offset, universityDeptID)
	if err != nil {
		return nil, err
	}

	squadIDs := make([]uuid.UUID, len(squads))
	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}
	participantsMap, _ := s.repo.GetParticipantsBySquadIDs(ctx, squadIDs)

	response := make([]dto.StudentSquadResponse, len(squads))
	for i, sq := range squads {
		position := ""
		if sq.OrganizerPosition != nil {
			position = *sq.OrganizerPosition
		}
		phone := ""
		if sq.OrganizerPhone != nil {
			phone = *sq.OrganizerPhone
		}
		enterpriseName := ""
		if sq.OrganizerEnterpriseName != nil {
			enterpriseName = *sq.OrganizerEnterpriseName
		}
		response[i] = dto.StudentSquadResponse{
			ID: sq.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:       sq.OrganizerID.String(),
				Name:     sq.OrganizerName,
				Role:     sq.OrganizerRole,
				Position: position,
				Phone:    phone,
				EnterpriseName: enterpriseName,
			},
			Title:           sq.Title,
			Description:     sq.Description,
			Profile:         sq.Profile,
			MaxParticipants: sq.MaxParticipants,
			CurrentCount:    sq.CurrentCount,
			Status:          sq.StatusName,
			StatusID:        sq.StatusID,
			Participants:    mapParticipants(participantsMap[sq.ID]),
			CreatedAt:       sq.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       sq.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      sq.ApprovedByName,
		}
	}

	return &dto.AllSquadsResponse{
		Squads: response,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (s *StudentSquadService) GetByID(ctx context.Context, id uuid.UUID) (*dto.StudentSquadDetailResponse, error) {
	squad, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if squad == nil {
		return nil, nil
	}

	participantIDs := make([]string, len(squad.Participants))
	for i, p := range squad.Participants {
		participantIDs[i] = p.String()
	}

	return &dto.StudentSquadDetailResponse{
		StudentSquadResponse: dto.StudentSquadResponse{
			ID: squad.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:   squad.OrganizerID.String(),
				Name: squad.OrganizerName,
				Role: squad.OrganizerRole,
				Position: *squad.OrganizerPosition,
				Phone:   *squad.OrganizerPhone,
			},
			Title:           squad.Title,
			Description:     squad.Description,
			Profile:         squad.Profile,
			MaxParticipants: squad.MaxParticipants,
			CurrentCount:    squad.CurrentCount,
			Status:          squad.StatusName,
			StatusID:        squad.StatusID,
			CreatedAt:       squad.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       squad.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      squad.ApprovedByName,
		},
		ParticipantIDs: participantIDs,
	}, nil
}

func (s *StudentSquadService) Create(ctx context.Context, organizerID uuid.UUID, req *dto.CreateStudentSquadRequest) (*dto.StudentSquadResponse, error) {
	statusName := "pending"
	if req.Role == "brsm" {
		statusName = "recruitment_open"
	}

	squad, err := s.repo.Create(ctx, organizerID, req.Title, req.Description, req.Profile, req.MaxParticipants, statusName)
	if err != nil {
		return nil, err
	}

	return &dto.StudentSquadResponse{
		ID: squad.ID.String(),
		Organizer: dto.SquadOrganizer{
			ID:   squad.OrganizerID.String(),
			Name: "",
			Role: "",
		},
		Title:           squad.Title,
		Description:     squad.Description,
		Profile:         squad.Profile,
		MaxParticipants: squad.MaxParticipants,
		CurrentCount:    0,
		Status:          statusName,
		StatusID:        squad.StatusID,
		CreatedAt:       squad.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:       squad.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *StudentSquadService) Update(ctx context.Context, id uuid.UUID, req *dto.UpdateStudentSquadRequest) error {
	return s.repo.Update(ctx, id, req.Title, req.Description, req.Profile, req.MaxParticipants, req.Status)
}

func (s *StudentSquadService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *StudentSquadService) Join(ctx context.Context, squadID, userID uuid.UUID) error {
	return s.repo.Join(ctx, squadID, userID)
}

func (s *StudentSquadService) Leave(ctx context.Context, squadID, userID uuid.UUID) error {
	return s.repo.Leave(ctx, squadID, userID)
}

func (s *StudentSquadService) IsUserParticipant(ctx context.Context, squadID, userID uuid.UUID) (bool, error) {
	return s.repo.IsUserParticipant(ctx, squadID, userID)
}

func (s *StudentSquadService) GetStatuses(ctx context.Context) ([]dto.SquadStatusResponse, error) {
	statuses, err := s.repo.GetStatuses(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.SquadStatusResponse, len(statuses))
	for i, st := range statuses {
		response[i] = dto.SquadStatusResponse{
			ID:          st.ID,
			Name:        st.Name,
			Description: st.Description,
		}
	}
	return response, nil
}

func (s *StudentSquadService) GetByOrganizer(ctx context.Context, organizerID uuid.UUID) ([]dto.StudentSquadResponse, error) {
	squads, err := s.repo.GetByOrganizer(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	squadIDs := make([]uuid.UUID, len(squads))
	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}
	participantsMap, _ := s.repo.GetParticipantsBySquadIDs(ctx, squadIDs)

	response := make([]dto.StudentSquadResponse, len(squads))
	for i, sq := range squads {
		position := ""
		if sq.OrganizerPosition != nil {
			position = *sq.OrganizerPosition
		}
		phone := ""
		if sq.OrganizerPhone != nil {
			phone = *sq.OrganizerPhone
		}
		enterpriseName := ""
		if sq.OrganizerEnterpriseName != nil {
			enterpriseName = *sq.OrganizerEnterpriseName
		}
		response[i] = dto.StudentSquadResponse{
			ID: sq.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:       sq.OrganizerID.String(),
				Name:     sq.OrganizerName,
				Role:     sq.OrganizerRole,
				Position: position,
				Phone:    phone,
				EnterpriseName: enterpriseName,
			},
			Title:           sq.Title,
			Description:     sq.Description,
			Profile:         sq.Profile,
			MaxParticipants: sq.MaxParticipants,
			CurrentCount:    sq.CurrentCount,
			Status:          sq.StatusName,
			StatusID:        sq.StatusID,
			Participants:    mapParticipants(participantsMap[sq.ID]),
			CreatedAt:       sq.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       sq.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      sq.ApprovedByName,
		}
	}
	return response, nil
}

func (s *StudentSquadService) GetByParticipant(ctx context.Context, userID uuid.UUID) ([]dto.StudentSquadResponse, error) {
	squads, err := s.repo.GetByParticipant(ctx, userID)
	if err != nil {
		return nil, err
	}

	squadIDs := make([]uuid.UUID, len(squads))
	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}
	participantsMap, _ := s.repo.GetParticipantsBySquadIDs(ctx, squadIDs)

	response := make([]dto.StudentSquadResponse, len(squads))
	for i, sq := range squads {
		position := ""
		if sq.OrganizerPosition != nil {
			position = *sq.OrganizerPosition
		}
		phone := ""
		if sq.OrganizerPhone != nil {
			phone = *sq.OrganizerPhone
		}
		enterpriseName := ""
		if sq.OrganizerEnterpriseName != nil {
			enterpriseName = *sq.OrganizerEnterpriseName
		}
		response[i] = dto.StudentSquadResponse{
			ID: sq.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:       sq.OrganizerID.String(),
				Name:     sq.OrganizerName,
				Role:     sq.OrganizerRole,
				Position: position,
				Phone:    phone,
				EnterpriseName: enterpriseName,
			},
			Title:           sq.Title,
			Description:     sq.Description,
			Profile:         sq.Profile,
			MaxParticipants: sq.MaxParticipants,
			CurrentCount:    sq.CurrentCount,
			Status:          sq.StatusName,
			StatusID:        sq.StatusID,
			Participants:    mapParticipants(participantsMap[sq.ID]),
			CreatedAt:       sq.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       sq.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      sq.ApprovedByName,
		}
	}
	return response, nil
}

func (s *StudentSquadService) GetByUniversity(ctx context.Context, universityDeptID uuid.UUID) ([]dto.StudentSquadResponse, error) {
	squads, err := s.repo.GetByUniversity(ctx, universityDeptID)
	if err != nil {
		return nil, err
	}

	squadIDs := make([]uuid.UUID, len(squads))
	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}
	participantsMap, _ := s.repo.GetParticipantsBySquadIDs(ctx, squadIDs)

	response := make([]dto.StudentSquadResponse, len(squads))
	for i, sq := range squads {
		position := ""
		if sq.OrganizerPosition != nil {
			position = *sq.OrganizerPosition
		}
		phone := ""
		if sq.OrganizerPhone != nil {
			phone = *sq.OrganizerPhone
		}
		enterpriseName := ""
		if sq.OrganizerEnterpriseName != nil {
			enterpriseName = *sq.OrganizerEnterpriseName
		}
		response[i] = dto.StudentSquadResponse{
			ID: sq.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:       sq.OrganizerID.String(),
				Name:     sq.OrganizerName,
				Role:     sq.OrganizerRole,
				Position: position,
				Phone:    phone,
				EnterpriseName: enterpriseName,
			},
			Title:           sq.Title,
			Description:     sq.Description,
			Profile:         sq.Profile,
			MaxParticipants: sq.MaxParticipants,
			CurrentCount:    sq.CurrentCount,
			Status:          sq.StatusName,
			StatusID:        sq.StatusID,
			Participants:    mapParticipants(participantsMap[sq.ID]),
			CreatedAt:       sq.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       sq.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      sq.ApprovedByName,
		}
	}
	return response, nil
}

func (s *StudentSquadService) GetByApprover(ctx context.Context, approverID uuid.UUID) ([]dto.StudentSquadResponse, error) {
	squads, err := s.repo.GetByApprover(ctx, approverID)
	if err != nil {
		return nil, err
	}

	squadIDs := make([]uuid.UUID, len(squads))
	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}
	participantsMap, _ := s.repo.GetParticipantsBySquadIDs(ctx, squadIDs)

	response := make([]dto.StudentSquadResponse, len(squads))
	for i, sq := range squads {
		position := ""
		if sq.OrganizerPosition != nil {
			position = *sq.OrganizerPosition
		}
		phone := ""
		if sq.OrganizerPhone != nil {
			phone = *sq.OrganizerPhone
		}
		enterpriseName := ""
		if sq.OrganizerEnterpriseName != nil {
			enterpriseName = *sq.OrganizerEnterpriseName
		}
		response[i] = dto.StudentSquadResponse{
			ID: sq.ID.String(),
			Organizer: dto.SquadOrganizer{
				ID:       sq.OrganizerID.String(),
				Name:     sq.OrganizerName,
				Role:     sq.OrganizerRole,
				Position: position,
				Phone:    phone,
				EnterpriseName: enterpriseName,
			},
			Title:           sq.Title,
			Description:     sq.Description,
			Profile:         sq.Profile,
			MaxParticipants: sq.MaxParticipants,
			CurrentCount:    sq.CurrentCount,
			Status:          sq.StatusName,
			StatusID:        sq.StatusID,
			Participants:    mapParticipants(participantsMap[sq.ID]),
			CreatedAt:       sq.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:       sq.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			ApprovedBy:      sq.ApprovedByName,
		}
	}
	return response, nil
}

func (s *StudentSquadService) Approve(ctx context.Context, squadID uuid.UUID, approved bool, approvedByUserID uuid.UUID) error {
	var statusName string
	if approved {
		statusName = "recruitment_open"
	} else {
		statusName = "rejected"
	}
	return s.repo.UpdateStatus(ctx, squadID, statusName, approvedByUserID)
}
