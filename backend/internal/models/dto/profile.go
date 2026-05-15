package dto

import "github.com/google/uuid"

type BrsmProfileRequest struct {
	FullName string  `json:"full_name" binding:"required"`
	Subrole  string  `json:"subrole"`
	Position *string `json:"position"`
	Phone    *string `json:"phone"`
}

type StudentProfileRequest struct {
	FullName               string     `json:"full_name" binding:"required"`
	Specialty              *string    `json:"specialty"`
	Grade                  *float64   `json:"grade" binding:"gte=1,lte=10"`
	UniversityDepartmentID *uuid.UUID `json:"university_department_id"`
	Position               *string    `json:"position"`
	Phone                  *string    `json:"phone"`
}

type UniversityProfileRequest struct {
	FullName               string     `json:"full_name" binding:"required"`
	Subrole                string     `json:"subrole"`
	UniversityDepartmentID *uuid.UUID `json:"university_department_id"`
	Position               *string    `json:"position"`
	Phone                  *string    `json:"phone"`
}

type CustomerProfileRequest struct {
	FullName     string     `json:"full_name" binding:"required"`
	EnterpriseID *uuid.UUID `json:"enterprise_id"`
	Position     *string    `json:"position"`
	Phone        *string    `json:"phone"`
}

type UpdateSubroleRequest struct {
	Role                   string   `json:"role" binding:"required"`
	FullName               *string  `json:"full_name"`
	Subrole                *string  `json:"subrole"`
	Specialty              *string  `json:"specialty"`
	Grade                  *float64 `json:"grade"`
	UniversityDepartmentID *string  `json:"university_department_id"`
	Position               *string  `json:"position"`
	Phone                  *string  `json:"phone"`
}

type ProfileResponse struct {
	UserID    string  `json:"user_id"`
	Email     string  `json:"email"`
	Login     string  `json:"login"`
	Role      string  `json:"role"`
	Avatar    *string `json:"avatar"`
	Profile   any     `json:"profile"`
	UpdatedAt string  `json:"updated_at"`
}

type BrsmProfileData struct {
	FullName string  `json:"full_name"`
	Subrole  *string `json:"subrole"`
	Position *string `json:"position"`
	Phone    *string `json:"phone"`
}

type UniversityDepartmentInfo struct {
	ID             string  `json:"id"`
	UniversityName string  `json:"university_name"`
	DepartmentName string  `json:"department_name"`
	Address        *string `json:"address"`
}

type StudentProfileData struct {
	FullName   string                    `json:"full_name"`
	Specialty  *string                   `json:"specialty"`
	Grade      *float64                  `json:"grade"`
	Position   *string                   `json:"position"`
	Phone      *string                   `json:"phone"`
	University *UniversityDepartmentInfo `json:"university,omitempty"`
}

type UniversityProfileData struct {
	FullName   string                    `json:"full_name"`
	Subrole    *string                   `json:"subrole"`
	Position   *string                   `json:"position"`
	Phone      *string                   `json:"phone"`
	University *UniversityDepartmentInfo `json:"university,omitempty"`
}

type EnterpriseInfo struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address,omitempty"`
}

type CustomerProfileData struct {
	FullName   string          `json:"full_name"`
	Enterprise *EnterpriseInfo `json:"enterprise,omitempty"`
	Position   *string         `json:"position"`
	Phone      *string         `json:"phone"`
}

type UserProfileData struct {
	FullName   string                    `json:"full_name"`
	Phone      *string                   `json:"phone"`
	Position   *string                   `json:"position"`
	University *UniversityDepartmentInfo `json:"university,omitempty"`
	Grade      *float64                  `json:"grade,omitempty"`
}

type AllProfilesResponse struct {
	Profiles []ProfileResponse `json:"profiles"`
	Total    int               `json:"total"`
	Limit    int               `json:"limit"`
	Offset   int               `json:"offset"`
}

type UserSearchResult struct {
	UserID    string  `json:"user_id"`
	Email     string  `json:"email"`
	Login     string  `json:"login"`
	Role      string  `json:"role"`
	Avatar    *string `json:"avatar"`
	FullName  string  `json:"full_name"`
	UpdatedAt string  `json:"updated_at"`
}

type UserSearchResponse struct {
	Profiles []UserSearchResult `json:"profiles"`
	Total    int                `json:"total"`
}

type UniversityStatsResponse struct {
	UniversityName   string `json:"university_name"`
	DepartmentName   string `json:"department_name"`
	TotalStudents    int    `json:"total_students"`
	StudentsInSquads int    `json:"students_in_squads"`
	StudentsEmployed int    `json:"students_employed"`
	TotalSquads      int    `json:"total_squads"`
	JobInvitations   int    `json:"job_invitations"`
}

type StudentSquadInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	JoinedAt    string  `json:"joined_at"`
}

type WorkInfo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Company   string `json:"company"`
	Position  string `json:"position"`
	Status    string `json:"status"`
	StartedAt string `json:"started_at"`
}

type StudentEmploymentInfo struct {
	Student       ProfileResponse    `json:"student"`
	StudentSquads []StudentSquadInfo `json:"student_squads"`
	Works         []WorkInfo         `json:"works"`
}

type StudentEmploymentListResponse struct {
	Students []StudentEmploymentInfo `json:"students"`
	Total    int                     `json:"total"`
	Limit    int                     `json:"limit"`
	Offset   int                     `json:"offset"`
}
