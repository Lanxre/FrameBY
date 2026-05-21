package db

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID           uuid.UUID
	Email        string
	Login        string
	PasswordHash string
	Role         string
	Avatar       *string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	FullName       string
	Subrole        string    // только для BRSM
	EnterpriseID   *uuid.UUID // только для Customer
	UniversityName string    // только для University
	DepartmentName string    // только для University
}
