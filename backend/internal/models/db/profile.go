package db

import (
	"time"

	"github.com/google/uuid"
)

type BrsmProfileEntity struct {
	UserID    uuid.UUID
	FullName  string
	Subrole   string
	Position  *string
	Phone     *string
	UpdatedAt time.Time
}

type StudentProfileEntity struct {
	UserID                 uuid.UUID
	FullName               string
	Position               *string
	Phone                  *string
	Specialty              *string
	SpecialtyID            *uuid.UUID
	Grade                  *float64
	UniversityDepartmentID *uuid.UUID
	UpdatedAt              time.Time
}

type UniversityProfileEntity struct {
	UserID                 uuid.UUID
	FullName               string
	Subrole                string
	Position               *string
	Phone                  *string
	UniversityDepartmentID *uuid.UUID
	UpdatedAt              time.Time
}

type CustomerProfileEntity struct {
	UserID       uuid.UUID
	FullName     string
	Subrole      *string
	EnterpriseID *uuid.UUID
	Position     *string
	Phone        *string
	UpdatedAt    time.Time
}

type AllProfilesResult struct {
	UserID                     uuid.UUID  `json:"user_id"`
	Email                      string     `json:"email"`
	Login                      string     `json:"login"`
	Role                       string     `json:"role"`
	Avatar                     *string    `json:"avatar"`
	FullName                   string     `json:"full_name"`
	Subrole                    *string    `json:"subrole"`
	Position                   *string    `json:"position"`
	Phone                      *string    `json:"phone"`
	Specialty                  *string    `json:"specialty"`
	Grade                      *float64   `json:"grade"`
	EnterpriseID               *uuid.UUID `json:"enterprise_id"`
	StudentUniversityDeptID    *uuid.UUID `json:"student_university_dept_id"`
	UniversityUniversityDeptID *uuid.UUID `json:"university_university_dept_id"`
	UpdatedAt                  string     `json:"updated_at"`
}
