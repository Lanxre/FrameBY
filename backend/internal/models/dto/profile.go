package dto

import "github.com/google/uuid"

type BrsmProfileRequest struct {
	FullName string  `json:"full_name" binding:"required"`
	Subrole  string  `json:"subrole"`
	Position *string `json:"position"`
	Phone    *string `json:"phone"`
}

type StudentProfileRequest struct {
	FullName  string   `json:"full_name" binding:"required"`
	Faculty   string   `json:"faculty" binding:"required"`
	Specialty string   `json:"specialty" binding:"required"`
	Grade     *float64 `json:"grade" binding:"required,gte=1,lte=10"`
	Position  *string  `json:"position"`
	Phone     *string  `json:"phone"`
}

type UniversityProfileRequest struct {
	FullName   string  `json:"full_name" binding:"required"`
	Subrole    string  `json:"subrole"`
	Faculty    string  `json:"faculty"`
	Department string  `json:"department" binding:"required"`
	Position   *string `json:"position"`
	Phone      *string `json:"phone"`
}

type CustomerProfileRequest struct {
	FullName     string     `json:"full_name" binding:"required"`
	EnterpriseID *uuid.UUID `json:"enterprise_id"`
	Position     *string    `json:"position"`
	Phone        *string    `json:"phone"`
}

type UpdateSubroleRequest struct {
	Role       string   `json:"role" binding:"required"`
	FullName   *string  `json:"full_name"`
	Subrole    *string  `json:"subrole"`
	Faculty    *string  `json:"faculty"`
	Specialty  *string  `json:"specialty"`
	Grade      *float64 `json:"grade"`
	Department *string  `json:"department"`
	Position   *string  `json:"position"`
}

type ProfileResponse struct {
	UserID       string   `json:"user_id"`
	Email        string   `json:"email"`
	Login        string   `json:"login"`
	Role         string   `json:"role"`
	Avatar       *string  `json:"avatar"`
	FullName     string   `json:"full_name"`
	Subrole      *string  `json:"subrole"`
	Position     *string  `json:"position"`
	Phone        *string  `json:"phone"`
	Faculty      *string  `json:"faculty"`
	Specialty    *string  `json:"specialty"`
	Grade        *float64 `json:"grade"`
	Department   *string  `json:"department"`
	EnterpriseID *string  `json:"enterprise_id"`
	UpdatedAt    string   `json:"updated_at"`
}

type AllProfilesResponse struct {
	Profiles []ProfileResponse `json:"profiles"`
	Total    int               `json:"total"`
	Limit    int               `json:"limit"`
	Offset   int               `json:"offset"`
}
