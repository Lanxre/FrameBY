
package services

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/google/uuid"

	"github.com/lanxre/frameby/internal/frender"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/repositories"
	"github.com/lanxre/frameby/internal/utils"
)

var ErrUnsupportedFormat = errors.New("unsupported export format")

type BytesRenderer interface {
	Bytes() ([]byte, error)
	Ext() string
}

type StudentSquadExportService struct {
	squadRepo               *repositories.StudentSquadRepository
	profileRepo             *repositories.ProfileRepository
	universityDepartmentRepo *repositories.UniversityDepartmentRepository
}

func NewStudentSquadExportService(
	squadRepo *repositories.StudentSquadRepository,
	profileRepo *repositories.ProfileRepository,
	universityDepartmentRepo *repositories.UniversityDepartmentRepository,
) *StudentSquadExportService {
	return &StudentSquadExportService{
		squadRepo:               squadRepo,
		profileRepo:             profileRepo,
		universityDepartmentRepo: universityDepartmentRepo,
	}
}

func (s *StudentSquadExportService) ExportSquad(
	ctx context.Context,
	squadID uuid.UUID,
	format string,
) ([]byte, string, error) {

	squad, err := s.squadRepo.GetByID(ctx, squadID)
	if err != nil {
		return nil, "", fmt.Errorf("get squad: %w", err)
	}

	if squad == nil {
		return nil, "", fmt.Errorf("squad not found")
	}

	participantsMap, err := s.squadRepo.GetParticipantsBySquadIDs(
		ctx,
		[]uuid.UUID{squadID},
	)
	if err != nil {
		return nil, "", fmt.Errorf("get participants: %w", err)
	}

	renderer, err := buildSingleRenderer(
		format,
		squad,
		participantsMap[squadID],
	)
	if err != nil {
		return nil, "", err
	}

	data, err := renderer.Bytes()
	if err != nil {
		return nil, "", fmt.Errorf("render export: %w", err)
	}

	filename := utils.SanitizeFilename(squad.Title) + renderer.Ext()

	return data, filename, nil
}

func (s *StudentSquadExportService) ExportAllSquads(
	ctx context.Context,
	userID uuid.UUID,
	role string,
	format string,
	myOnly bool,
) ([]byte, string, error) {

	squads, err := s.loadSquads(ctx, userID, role, myOnly)
	if err != nil {
		return nil, "", err
	}

	sort.Slice(squads, func(i, j int) bool {
		return squadPriority(squads[i]) < squadPriority(squads[j])
	})

	squadIDs := make([]uuid.UUID, len(squads))

	for i, sq := range squads {
		squadIDs[i] = sq.ID
	}

	participantsMap, err := s.squadRepo.GetParticipantsBySquadIDs(
		ctx,
		squadIDs,
	)
	if err != nil {
		return nil, "", fmt.Errorf("get participants: %w", err)
	}

	requesterLabel := s.buildRequesterLabel(ctx, userID, role)

	renderer, err := buildMultiRenderer(
		format,
		squads,
		participantsMap,
		requesterLabel,
	)
	if err != nil {
		return nil, "", err
	}

	data, err := renderer.Bytes()
	if err != nil {
		return nil, "", fmt.Errorf("render export: %w", err)
	}

	filename := "student_squads" + renderer.Ext()

	return data, filename, nil
}

func (s *StudentSquadExportService) loadSquads(
	ctx context.Context,
	userID uuid.UUID,
	role string,
	myOnly bool,
) ([]db.StudentSquadWithDetails, error) {

	if !myOnly {
		squads, err := s.squadRepo.GetAllSquads(ctx)
		if err != nil {
			return nil, fmt.Errorf("get squads: %w", err)
		}

		return squads, nil
	}

	switch role {

	case "student":
		squads, err := s.squadRepo.GetByParticipant(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("get participant squads: %w", err)
		}

		return squads, nil

	case "university":
		squads, err := s.squadRepo.GetByApprover(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("get approver squads: %w", err)
		}

		return squads, nil

	default:
		squads, err := s.squadRepo.GetByOrganizer(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("get organizer squads: %w", err)
		}

		return squads, nil
	}
}

func (s *StudentSquadExportService) buildRequesterLabel(
	ctx context.Context,
	userID uuid.UUID,
	role string,
) string {

	switch role {

	case "university":
		return s.buildUniversityLabel(ctx, userID)

	case "student":
		return s.buildStudentLabel(ctx, userID)

	default:
		return s.buildBRSMLabel(ctx, userID)
	}
}

func (s *StudentSquadExportService) buildUniversityLabel(
	ctx context.Context,
	userID uuid.UUID,
) string {

	profile, err := s.profileRepo.GetUniversityProfile(ctx, userID)
	if err != nil || profile == nil {
		return ""
	}

	if profile.UniversityDepartmentID != nil {

		dept, err := s.universityDepartmentRepo.GetByID(
			ctx,
			*profile.UniversityDepartmentID,
		)

		if err == nil &&
			dept != nil &&
			dept.University != nil {

			return fmt.Sprintf(
				"Университет: %s • %s",
				dept.University.Name,
				profile.FullName,
			)
		}
	}

	return fmt.Sprintf(
		"Университет • %s",
		profile.FullName,
	)
}

func (s *StudentSquadExportService) buildStudentLabel(
	ctx context.Context,
	userID uuid.UUID,
) string {

	profile, err := s.profileRepo.GetStudentProfile(ctx, userID)
	if err != nil || profile == nil {
		return ""
	}

	return fmt.Sprintf(
		"Студент • %s",
		profile.FullName,
	)
}

func (s *StudentSquadExportService) buildBRSMLabel(
	ctx context.Context,
	userID uuid.UUID,
) string {

	profile, err := s.profileRepo.GetBrsmProfile(ctx, userID)
	if err != nil || profile == nil {
		return ""
	}

	return fmt.Sprintf(
		"БРСМ • %s",
		profile.FullName,
	)
}

func squadPriority(s db.StudentSquadWithDetails) int {
	if s.CurrentCount > 0 {
		return 0
	}

	return 1
}

func buildSingleRenderer(
	format string,
	squad *db.StudentSquadWithParticipants,
	participants []db.SquadParticipantInfo,
) (BytesRenderer, error) {

	switch format {

	case "md":
		return frender.NewSquadMarkdown(
			squad,
			participants,
		), nil

	case "html":
		return frender.NewSquadHTML(
			squad,
			participants,
		), nil

	case "pdf":
		return frender.NewSquadPDF(
			squad,
			participants,
		), nil

	default:
		return nil, fmt.Errorf(
			"%w: %s",
			ErrUnsupportedFormat,
			format,
		)
	}
}

func buildMultiRenderer(
	format string,
	squads []db.StudentSquadWithDetails,
	participantsMap map[uuid.UUID][]db.SquadParticipantInfo,
	requesterLabel string,
) (BytesRenderer, error) {

	switch format {

	case "md":
		return frender.NewSquadsMarkdown(
			squads,
			participantsMap,
			requesterLabel,
		), nil

	case "html":
		return frender.NewSquadsHTML(
			squads,
			participantsMap,
			requesterLabel,
		), nil

	case "pdf":
		return frender.NewSquadsPDF(
			squads,
			participantsMap,
			requesterLabel,
		), nil

	default:
		return nil, fmt.Errorf(
			"%w: %s",
			ErrUnsupportedFormat,
			format,
		)
	}
}
