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
	repo               *repositories.ProfileRepository
	userRepo           *repositories.UserRepository
	universityDeptRepo *repositories.UniversityDepartmentRepository
	enterpriseRepo     *repositories.EnterpriseRepository
}

func NewProfileService(repo *repositories.ProfileRepository, userRepo *repositories.UserRepository, universityDeptRepo *repositories.UniversityDepartmentRepository, enterpriseRepo *repositories.EnterpriseRepository) *ProfileService {
	return &ProfileService{
		repo:               repo,
		userRepo:           userRepo,
		universityDeptRepo: universityDeptRepo,
		enterpriseRepo:     enterpriseRepo,
	}
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
	return s.repo.CreateStudentProfile(ctx, userID, req.FullName, req.Specialty, req.Grade, req.Position, req.Phone, req.UniversityDepartmentID)
}

func (s *ProfileService) UpdateStudentProfile(ctx context.Context, userID uuid.UUID, req *dto.StudentProfileRequest) error {
	return s.repo.UpdateStudentProfile(ctx, userID, req.FullName, req.Specialty, req.Grade, req.Position, req.Phone, req.UniversityDepartmentID)
}

func (s *ProfileService) DeleteStudentProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteStudentProfile(ctx, userID)
}

func (s *ProfileService) GetUniversityProfile(ctx context.Context, userID uuid.UUID) (*db.UniversityProfileEntity, error) {
	return s.repo.GetUniversityProfile(ctx, userID)
}

func (s *ProfileService) CreateUniversityProfile(ctx context.Context, userID uuid.UUID, req *dto.UniversityProfileRequest) error {
	return s.repo.CreateUniversityProfile(ctx, userID, req.FullName, req.Subrole, req.Position, req.Phone, req.UniversityDepartmentID)
}

func (s *ProfileService) UpdateUniversityProfile(ctx context.Context, userID uuid.UUID, req *dto.UniversityProfileRequest) error {
	return s.repo.UpdateUniversityProfile(ctx, userID, req.FullName, req.Subrole, req.Position, req.Phone, req.UniversityDepartmentID)
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

func (s *ProfileService) GetUniversityStats(ctx context.Context, universityDeptID uuid.UUID) (*dto.UniversityStatsResponse, error) {
	dept, err := s.universityDeptRepo.GetByID(ctx, universityDeptID)
	if err != nil || dept == nil || dept.University == nil || dept.Department == nil {
		return nil, errors.New("university department not found")
	}

	stats, err := s.repo.GetUniversityStats(ctx, universityDeptID)
	if err != nil {
		return nil, err
	}

	return &dto.UniversityStatsResponse{
		UniversityName:   dept.University.Name,
		DepartmentName:   dept.Department.Name,
		TotalStudents:    stats.TotalStudents,
		StudentsInSquads: stats.StudentsInSquads,
		StudentsEmployed: stats.StudentsEmployed,
		TotalSquads:      stats.TotalSquads,
		JobInvitations:   stats.JobInvitations,
	}, nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteAllProfiles(ctx, userID)
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
			UserID:    p.UserID.String(),
			Email:     p.Email,
			Login:     p.Login,
			Role:      p.Role,
			Avatar:    p.Avatar,
			UpdatedAt: p.UpdatedAt,
		}

		switch p.Role {
		case string(types.RoleBRSM):
			profilesResponse[i].Profile = dto.BrsmProfileData{
				FullName: p.FullName,
				Subrole:  p.Subrole,
				Position: p.Position,
				Phone:    p.Phone,
			}
		case string(types.RoleStudent):
			var universityInfo *dto.UniversityDepartmentInfo
			if p.StudentUniversityDeptID != nil {
				dept, _ := s.universityDeptRepo.GetByID(ctx, *p.StudentUniversityDeptID)
				if dept != nil && dept.University != nil && dept.Department != nil {
					universityInfo = &dto.UniversityDepartmentInfo{
						ID:             dept.ID.String(),
						UniversityName: dept.University.Name,
						DepartmentName: dept.Department.Name,
						Address:        dept.Address,
					}
				}
			}
			profilesResponse[i].Profile = dto.StudentProfileData{
				FullName:   p.FullName,
				Specialty:  p.Specialty,
				Grade:      p.Grade,
				Position:   p.Position,
				Phone:      p.Phone,
				University: universityInfo,
			}
		case string(types.RoleUniversity):
			var universityInfo *dto.UniversityDepartmentInfo
			if p.UniversityUniversityDeptID != nil {
				dept, _ := s.universityDeptRepo.GetByID(ctx, *p.UniversityUniversityDeptID)
				if dept != nil && dept.University != nil && dept.Department != nil {
					universityInfo = &dto.UniversityDepartmentInfo{
						ID:             dept.ID.String(),
						UniversityName: dept.University.Name,
						DepartmentName: dept.Department.Name,
						Address:        dept.Address,
					}
				}
			}
			profilesResponse[i].Profile = dto.UniversityProfileData{
				FullName:   p.FullName,
				Subrole:    p.Subrole,
				Position:   p.Position,
				Phone:      p.Phone,
				University: universityInfo,
			}
		case string(types.RoleCustomer):
			var enterpriseInfo *dto.EnterpriseInfo
			if p.EnterpriseID != nil {
				enterprise, _ := s.enterpriseRepo.GetByID(ctx, *p.EnterpriseID)
				if enterprise != nil {
					enterpriseInfo = &dto.EnterpriseInfo{
						ID:      enterprise.ID.String(),
						Name:    enterprise.Name,
						Address: enterprise.Address,
					}
				}
			}
			profilesResponse[i].Profile = dto.CustomerProfileData{
				FullName:   p.FullName,
				Enterprise: enterpriseInfo,
				Position:   p.Position,
				Phone:      p.Phone,
			}
		default:
			var universityInfo *dto.UniversityDepartmentInfo
			if p.StudentUniversityDeptID != nil {
				dept, _ := s.universityDeptRepo.GetByID(ctx, *p.StudentUniversityDeptID)
				if dept != nil && dept.University != nil && dept.Department != nil {
					universityInfo = &dto.UniversityDepartmentInfo{
						ID:             dept.ID.String(),
						UniversityName: dept.University.Name,
						DepartmentName: dept.Department.Name,
						Address:        dept.Address,
					}
				}
			} else if p.UniversityUniversityDeptID != nil {
				dept, _ := s.universityDeptRepo.GetByID(ctx, *p.UniversityUniversityDeptID)
				if dept != nil && dept.University != nil && dept.Department != nil {
					universityInfo = &dto.UniversityDepartmentInfo{
						ID:             dept.ID.String(),
						UniversityName: dept.University.Name,
						DepartmentName: dept.Department.Name,
						Address:        dept.Address,
					}
				}
			}

			profilesResponse[i].Profile = dto.UserProfileData{
				FullName:   p.FullName,
				Phone:      p.Phone,
				Position:   p.Position,
				University: universityInfo,
				Grade:      p.Grade,
			}
		}
	}

	return &dto.AllProfilesResponse{
		Profiles: profilesResponse,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

type UniversityDepartmentInfo struct {
	UniversityName string
	DepartmentName string
}

func (s *ProfileService) GetUniversityDepartmentInfo(ctx context.Context, deptID uuid.UUID) (*UniversityDepartmentInfo, error) {
	dept, err := s.universityDeptRepo.GetByID(ctx, deptID)
	if err != nil || dept == nil {
		return nil, err
	}
	if dept.University == nil || dept.Department == nil {
		return nil, errors.New("university or department not found")
	}
	return &UniversityDepartmentInfo{
		UniversityName: dept.University.Name,
		DepartmentName: dept.Department.Name,
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
		return s.repo.UpdateUniversityProfile(ctx, userID, profile.FullName, subrole, profile.Position, profile.Phone, profile.UniversityDepartmentID)
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
		position := req.Position
		phone := req.Phone
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if req.Position == nil {
				position = profile.Position
			}
			if req.Phone == nil {
				phone = profile.Phone
			}
			return s.repo.UpdateBrsmProfile(ctx, userID, fullName, ptrToStr(req.Subrole), position, phone)
		}
		if fullName == "" {
			return errors.New("full_name обязательно для заполнения")
		}
		return s.repo.CreateBrsmProfile(ctx, userID, fullName, ptrToStr(req.Subrole), position, phone)

	case string(types.RoleStudent):
		profile, _ := s.repo.GetStudentProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		specialty := req.Specialty
		grade := req.Grade
		position := req.Position
		phone := req.Phone
		var univDeptID *uuid.UUID
		if req.UniversityDepartmentID != nil {
			id, err := uuid.Parse(*req.UniversityDepartmentID)
			if err == nil {
				univDeptID = &id
			}
		}
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if specialty == nil {
				specialty = profile.Specialty
			}
			if grade == nil {
				grade = profile.Grade
			}
			if req.Position == nil {
				position = profile.Position
			}
			if req.Phone == nil {
				phone = profile.Phone
			}
			if univDeptID == nil {
				univDeptID = profile.UniversityDepartmentID
			}
			return s.repo.UpdateStudentProfile(ctx, userID, fullName, specialty, grade, position, phone, univDeptID)
		}
		if fullName == "" {
			return errors.New("full_name обязательно для заполнения")
		}
		return s.repo.CreateStudentProfile(ctx, userID, fullName, specialty, grade, position, phone, univDeptID)

	case string(types.RoleUniversity):
		profile, _ := s.repo.GetUniversityProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		position := req.Position
		phone := req.Phone
		var univDeptID *uuid.UUID
		if req.UniversityDepartmentID != nil {
			id, err := uuid.Parse(*req.UniversityDepartmentID)
			if err == nil {
				univDeptID = &id
			}
		}
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if req.Position != nil {
				position = req.Position
			} else {
				position = profile.Position
			}
			if req.Phone == nil {
				phone = profile.Phone
			}
			if univDeptID == nil {
				univDeptID = profile.UniversityDepartmentID
			}
			return s.repo.UpdateUniversityProfile(ctx, userID, fullName, ptrToStr(req.Subrole), position, phone, univDeptID)
		}
		if fullName == "" {
			return errors.New("full_name обязательно для заполнения")
		}
		return s.repo.CreateUniversityProfile(ctx, userID, fullName, ptrToStr(req.Subrole), position, phone, univDeptID)

	case string(types.RoleCustomer):
		profile, _ := s.repo.GetCustomerProfile(ctx, userID)
		fullName := ptrToStr(req.FullName)
		position := req.Position
		phone := req.Phone
		if profile != nil {
			if req.FullName != nil && *req.FullName != "" {
				fullName = *req.FullName
			} else {
				fullName = profile.FullName
			}
			if req.Position == nil {
				position = profile.Position
			}
			if req.Phone == nil {
				phone = profile.Phone
			}
			return s.repo.UpdateCustomerProfile(ctx, userID, fullName, profile.EnterpriseID, position, phone)
		}
		if fullName == "" {
			return errors.New("full_name обязательно для заполнения")
		}
		return s.repo.CreateCustomerProfile(ctx, userID, fullName, nil, position, phone)
	}

	return nil
}

func ptrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (s *ProfileService) SearchUsers(ctx context.Context, search string, limit int) (*dto.AllProfilesResponse, error) {
	return s.repo.SearchUsers(ctx, search, limit)
}

func (s *ProfileService) GetStudentsWithEmployment(ctx context.Context, universityDeptID uuid.UUID, limit, offset int) (*dto.StudentEmploymentListResponse, error) {
	return s.repo.GetStudentsWithEmployment(ctx, universityDeptID, limit, offset)
}
