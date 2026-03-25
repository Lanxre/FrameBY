package db

import (
	"time"

	"github.com/google/uuid"
)

type EmploymentStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmploymentParticipantStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmploymentRequest struct {
	ID                     uuid.UUID `json:"id"`
	EnterpriseID           uuid.UUID `json:"enterprise_id"`
	UniversityDepartmentID uuid.UUID `json:"university_department_id"`
	Title                  string    `json:"title"`
	Description            *string   `json:"description"`
	Requirements           *string   `json:"requirements"`
	Salary                 *string   `json:"salary"`
	Schedule               *string   `json:"schedule"`
	MaxParticipants        int       `json:"max_participants"`
	StatusID               int       `json:"status_id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type EmploymentParticipant struct {
	ID           uuid.UUID  `json:"id"`
	RequestID    uuid.UUID  `json:"request_id"`
	UserID       uuid.UUID  `json:"user_id"`
	StatusID     int        `json:"status_id"`
	AppliedAt    time.Time  `json:"applied_at"`
	ContractedAt *time.Time `json:"contracted_at"`
}

type EmploymentRequestWithDetails struct {
	ID                     uuid.UUID `json:"id"`
	EnterpriseID           uuid.UUID `json:"enterprise_id"`
	EnterpriseName         string    `json:"enterprise_name"`
	EnterpriseAddress      *string   `json:"enterprise_address"`
	UniversityDepartmentID uuid.UUID `json:"university_department_id"`
	UniversityName         string    `json:"university_name"`
	DepartmentName         string    `json:"department_name"`
	UniversityAddress      *string   `json:"university_address"`
	Title                  string    `json:"title"`
	Description            *string   `json:"description"`
	Requirements           *string   `json:"requirements"`
	Salary                 *string   `json:"salary"`
	Schedule               *string   `json:"schedule"`
	MaxParticipants        int       `json:"max_participants"`
	CurrentParticipants    int       `json:"current_participants"`
	StatusID               int       `json:"status_id"`
	StatusName             string    `json:"status_name"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type EmploymentRequestWithParticipants struct {
	EmploymentRequestWithDetails
	ParticipantUserIDs []uuid.UUID `json:"participant_user_ids"`
}

type EmploymentParticipantWithDetails struct {
	ID           uuid.UUID  `json:"id"`
	RequestID    uuid.UUID  `json:"request_id"`
	UserID       uuid.UUID  `json:"user_id"`
	StudentName  string     `json:"student_name"`
	StatusID     int        `json:"status_id"`
	StatusName   string     `json:"status_name"`
	AppliedAt    time.Time  `json:"applied_at"`
	ContractedAt *time.Time `json:"contracted_at"`
}
