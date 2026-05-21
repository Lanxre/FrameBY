package db

import (
	"time"

	"github.com/google/uuid"
)

type SquadStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type StudentSquad struct {
	ID              uuid.UUID `json:"id"`
	OrganizerID     uuid.UUID `json:"organizer_id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description"`
	Profile         *string   `json:"profile"`
	MaxParticipants int       `json:"max_participants"`
	StatusID        int       `json:"status_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type StudentSquadParticipant struct {
	ID       uuid.UUID `json:"id"`
	SquadID  uuid.UUID `json:"squad_id"`
	UserID   uuid.UUID `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}

type StudentSquadWithDetails struct {
	ID                uuid.UUID `json:"id"`
	OrganizerID       uuid.UUID `json:"organizer_id"`
	OrganizerName     string    `json:"organizer_name"`
	OrganizerRole     string    `json:"organizer_role"`
	OrganizerPosition *string   `json:"organizer_position"`
	OrganizerPhone    *string   `json:"organizer_phone"`
	Title             string    `json:"title"`
	Description       *string   `json:"description"`
	Profile           *string   `json:"profile"`
	MaxParticipants   int       `json:"max_participants"`
	CurrentCount      int       `json:"current_count"`
	StatusID          int       `json:"status_id"`
	StatusName        string    `json:"status_name"`
	CreatedAt         time.Time `json:"created_at"`
	Avatar            *string   `json:"avatar"`
	UpdatedAt         time.Time `json:"updated_at"`
	ApprovedByName          string    `json:"approved_by_name"`
	OrganizerEnterpriseName *string   `json:"organizer_enterprise_name"`
}

type SquadParticipantInfo struct {
	SquadID   uuid.UUID `json:"squad_id"`
	UserID    uuid.UUID `json:"user_id"`
	FullName  string    `json:"full_name"`
	Phone     *string   `json:"phone"`
	Specialty *string   `json:"specialty"`
	Grade     *float64  `json:"grade"`
	Avatar    *string   `json:"avatar"`
}

type StudentSquadWithParticipants struct {
	StudentSquadWithDetails
	Participants []uuid.UUID `json:"participants"`
}
