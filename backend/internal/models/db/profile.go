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
	UserID    uuid.UUID
	FullName  string
	Position  *string
	Phone     *string
	Faculty   *string
	Specialty *string
	Grade     *float64
	UpdatedAt time.Time
}

type UniversityProfileEntity struct {
	UserID     uuid.UUID
	FullName   string
	Subrole    string
	Faculty    *string
	Department *string
	Position   *string
	Phone      *string
	UpdatedAt  time.Time
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
	UserID       uuid.UUID  `json:"user_id"`
	Email        string     `json:"email"`
	Login        string     `json:"login"`
	Role         string     `json:"role"`
	Avatar       *string    `json:"avatar"`
	FullName     string     `json:"full_name"`
	Subrole      *string    `json:"subrole"`
	Position     *string    `json:"position"`
	Phone        *string    `json:"phone"`
	Faculty      *string    `json:"faculty"`
	Specialty    *string    `json:"specialty"`
	Grade        *float64   `json:"grade"`
	Department   *string    `json:"department"`
	EnterpriseID *uuid.UUID `json:"enterprise_id"`
	UpdatedAt    string     `json:"updated_at"`
}
