package db

import (
	"time"

	"github.com/google/uuid"
)

type University struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Department struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UniversityDepartmentEntity struct {
	ID           uuid.UUID   `json:"id"`
	UniversityID uuid.UUID   `json:"university_id"`
	DepartmentID uuid.UUID   `json:"department_id"`
	Address      *string     `json:"address"`
	CreatedAt    time.Time   `json:"created_at"`
	University   *University `json:"university,omitempty"`
	Department   *Department `json:"department,omitempty"`
}

type UniversityDepartmentWithDetails struct {
	ID             uuid.UUID `json:"id"`
	UniversityName string    `json:"university_name"`
	DepartmentName string    `json:"department_name"`
	Address        *string   `json:"address"`
}
