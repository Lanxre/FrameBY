package db

import (
	"time"

	"github.com/google/uuid"
)

type SpecialtyEntity struct {
	ID           uuid.UUID
	Name         string
	DepartmentID uuid.UUID
	CreatedAt    time.Time
}
