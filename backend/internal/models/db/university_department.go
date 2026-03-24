package db

import (
	"time"

	"github.com/google/uuid"
)

type UniversityDepartmentEntity struct {
	ID             uuid.UUID `json:"id"`
	UniversityName string    `json:"university_name"`
	DepartmentName string    `json:"department_name"`
	Address        *string   `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
}
