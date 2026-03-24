package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
	"github.com/lanxre/frameby/internal/types"
)

type UpdateProfileData struct {
	Role       string
	Subrole    *string
	Faculty    *string
	Specialty  *string
	Grade      *float64
	Department *string
	Position   *string
}

type ProfileService struct {
	repo     *repositories.ProfileRepository
	userRepo *repositories.UserRepository
}

func NewProfileService(repo *repositories.ProfileRepository, userRepo *repositories.UserRepository) *ProfileService {
	return &ProfileService{repo: repo, userRepo: userRepo}
}

func (s *ProfileService) GetBrsmProfile(ctx context.Context, userID uuid.UUID) (*db.BrsmProfileEntity, error) {
	return s.repo.GetBrsmProfile(ctx, userID)
}

func (s *ProfileService) CreateBrsmProfile(ctx context.Context, userID uuid.UUID, req *dto.BrsmProfileRequest) error {
	return s.repo.CreateBrsmProfile(ctx, userID, req.FullName, req.Subrole, req.Position, req.Phone)
}

func (s *ProfileService) UpdateBrsmProfile(ctx context.Context, userID uuid.UUID, req *dto.BrsmProfileRequest) error {
	return s.repo.UpdateBrsmProfile(ctx, userID, req.FullName, req.Subrole, req.Position, req.Phone)
}

func (s *ProfileService) DeleteBrsmProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteBrsmProfile(ctx, userID)
}

func (s *ProfileService) GetStudentProfile(ctx context.Context, userID uuid.UUID) (*db.StudentProfileEntity, error) {
	return s.repo.GetStudentProfile(ctx, userID)
}

func (s *ProfileService) CreateStudentProfile(ctx context.Context, userID uuid.UUID, req *dto.StudentProfileRequest) error {
	return s.repo.CreateStudentProfile(ctx, userID, req.FullName, req.Faculty, req.Specialty, req.Grade, req.Position, req.Phone)
}

func (s *ProfileService) UpdateStudentProfile(ctx context.Context, userID uuid.UUID, req *dto.StudentProfileRequest) error {
	return s.repo.UpdateStudentProfile(ctx, userID, req.FullName, req.Faculty, req.Specialty, req.Grade, req.Position, req.Phone)
}

func (s *ProfileService) DeleteStudentProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteStudentProfile(ctx, userID)
}

func (s *ProfileService) GetUniversityProfile(ctx context.Context, userID uuid.UUID) (*db.UniversityProfileEntity, error) {
	return s.repo.GetUniversityProfile(ctx, userID)
}

func (s *ProfileService) CreateUniversityProfile(ctx context.Context, userID uuid.UUID, req *dto.UniversityProfileRequest) error {
	return s.repo.CreateUniversityProfile(ctx, userID, req.FullName, req.Subrole, req.Faculty, req.Department, req.Position, req.Phone)
}

func (s *ProfileService) UpdateUniversityProfile(ctx context.Context, userID uuid.UUID, req *dto.UniversityProfileRequest) error {
	return s.repo.UpdateUniversityProfile(ctx, userID, req.FullName, req.Subrole, req.Faculty, req.Department, req.Position, req.Phone)
}

func (s *ProfileService) DeleteUniversityProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteUniversityProfile(ctx, userID)
}

func (s *ProfileService) GetCustomerProfile(ctx context.Context, userID uuid.UUID) (*db.CustomerProfileEntity, error) {
	return s.repo.GetCustomerProfile(ctx, userID)
}

func (s *ProfileService) CreateCustomerProfile(ctx context.Context, userID uuid.UUID, req *dto.CustomerProfileRequest) error {
	return s.repo.CreateCustomerProfile(ctx, userID, req.FullName, req.EnterpriseID, req.Position, req.Phone)
}

func (s *ProfileService) UpdateCustomerProfile(ctx context.Context, userID uuid.UUID, req *dto.CustomerProfileRequest) error {
	return s.repo.UpdateCustomerProfile(ctx, userID, req.FullName, req.EnterpriseID, req.Position, req.Phone)
}

func (s *ProfileService) DeleteCustomerProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteCustomerProfile(ctx, userID)
}

func (s *ProfileService) UpdateProfileByRole(ctx context.Context, userID uuid.UUID, role string, req interface{}) error {
	switch role {
	case string(types.RoleBRSM):
		r, ok := req.(*dto.BrsmProfileRequest)
		if !ok {
			return errors.New("invalid request type for brsm")
		}
		return s.UpdateBrsmProfile(ctx, userID, r)
	case string(types.RoleStudent):
		r, ok := req.(*dto.StudentProfileRequest)
		if !ok {
			return errors.New("invalid request type for student")
		}
		return s.UpdateStudentProfile(ctx, userID, r)
	case string(types.RoleUniversity):
		r, ok := req.(*dto.UniversityProfileRequest)
		if !ok {
			return errors.New("invalid request type for university")
		}
		return s.UpdateUniversityProfile(ctx, userID, r)
	case string(types.RoleCustomer):
		r, ok := req.(*dto.CustomerProfileRequest)
		if !ok {
			return errors.New("invalid request type for customer")
		}
		return s.UpdateCustomerProfile(ctx, userID, r)
	default:
		return errors.New("invalid role")
	}
}

func (s *ProfileService) GetAllProfiles(ctx context.Context, profileType, search string, limit, offset int) (*dto.AllProfilesResponse, error) {
	profiles, total, err := s.repo.GetAllProfiles(ctx, profileType, search, limit, offset)
	if err != nil {
		return nil, err
	}

	profilesResponse := make([]dto.ProfileResponse, len(profiles))
	for i, p := range profiles {
		profilesResponse[i] = dto.ProfileResponse{
			UserID:     p.UserID.String(),
			Email:      p.Email,
			Login:      p.Login,
			Role:       p.Role,
			Avatar:     p.Avatar,
			FullName:   p.FullName,
			Subrole:    p.Subrole,
			Position:   p.Position,
			Phone:      p.Phone,
			Faculty:    p.Faculty,
			Specialty:  p.Specialty,
			Grade:      p.Grade,
			Department: p.Department,
			UpdatedAt:  p.UpdatedAt,
		}
		if p.EnterpriseID != nil {
			eid := p.EnterpriseID.String()
			profilesResponse[i].EnterpriseID = &eid
		}
	}

	return &dto.AllProfilesResponse{
		Profiles: profilesResponse,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

func (s *ProfileService) UpdateSubrole(ctx context.Context, userID uuid.UUID, role, subrole string) error {
	switch role {
	case string(types.RoleBRSM):
		profile, err := s.repo.GetBrsmProfile(ctx, userID)
		if err != nil {
			return err
		}
		if profile == nil {
			return errors.New("profile not found")
		}
		return s.repo.UpdateBrsmProfile(ctx, userID, profile.FullName, subrole, profile.Position, profile.Phone)
	case string(types.RoleUniversity):
		profile, err := s.repo.GetUniversityProfile(ctx, userID)
		if err != nil {
			return err
		}
		if profile == nil {
			return errors.New("profile not found")
		}
		return s.repo.UpdateUniversityProfile(ctx, userID, profile.FullName, subrole, ptrToStr(profile.Faculty), ptrToStr(profile.Department), profile.Position, profile.Phone)
	default:
		return errors.New("subrole not applicable for this role")
	}
}

func (s *ProfileService) UpdateProfile(ctx context.Context, userID uuid.UUID, req *dto.UpdateSubroleRequest) error {
	existing, _ := s.userRepo.GetByID(ctx, userID)
	if existing == nil {
		return errors.New("user not found")
	}

	newRole := req.Role
	roleChanged := newRole != "" && newRole != existing.Role

	if roleChanged {
		if err := s.userRepo.UpdateRole(ctx, userID, newRole); err != nil {
			return errors.New("failed to update user role")
		}
	}

	roleToUpdate := newRole
	if roleToUpdate == "" {
		roleToUpdate = existing.Role
	}

	switch roleToUpdate {
	case string(types.RoleBRSM):
		profile, _ := s.repo.GetBrsmProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			position := req.Position
			phone := profile.Phone
			return s.repo.UpdateBrsmProfile(ctx, userID, fullName, ptrToStr(req.Subrole), position, phone)
		}
		return s.repo.CreateBrsmProfile(ctx, userID, fullName, ptrToStr(req.Subrole), req.Position, nil)

	case string(types.RoleStudent):
		profile, _ := s.repo.GetStudentProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		faculty := req.Faculty
		specialty := req.Specialty
		grade := req.Grade
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if faculty == nil {
				faculty = profile.Faculty
			}
			if specialty == nil {
				specialty = profile.Specialty
			}
			if grade == nil {
				grade = profile.Grade
			}
			position := req.Position
			phone := profile.Phone
			return s.repo.UpdateStudentProfile(ctx, userID, fullName, ptrToStr(faculty), ptrToStr(specialty), grade, position, phone)
		}
		return s.repo.CreateStudentProfile(ctx, userID, fullName, ptrToStr(faculty), ptrToStr(specialty), grade, req.Position, nil)

	case string(types.RoleUniversity):
		profile, _ := s.repo.GetUniversityProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		faculty := req.Faculty
		department := req.Department
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if faculty == nil {
				faculty = profile.Faculty
			}
			if department == nil {
				department = profile.Department
			}
			position := req.Position
			phone := profile.Phone
			return s.repo.UpdateUniversityProfile(ctx, userID, fullName, ptrToStr(req.Subrole), ptrToStr(faculty), ptrToStr(department), position, phone)
		}
		return s.repo.CreateUniversityProfile(ctx, userID, fullName, ptrToStr(req.Subrole), ptrToStr(faculty), ptrToStr(department), req.Position, nil)

	case string(types.RoleCustomer):
		profile, _ := s.repo.GetCustomerProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			position := req.Position
			phone := profile.Phone
			return s.repo.UpdateCustomerProfile(ctx, userID, fullName, profile.EnterpriseID, position, phone)
		}
		return s.repo.CreateCustomerProfile(ctx, userID, fullName, nil, req.Position, nil)
	}

	return nil
}

func ptrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
