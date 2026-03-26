package dto

type CreateStudentSquadRequest struct {
	Title           string  `json:"title" binding:"required"`
	Description     *string `json:"description"`
	Profile         *string `json:"profile"`
	MaxParticipants int     `json:"max_participants" binding:"required,gt=0"`
	Role            string  `json:"-"`
}

type UpdateStudentSquadRequest struct {
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	Profile         *string `json:"profile"`
	MaxParticipants *int    `json:"max_participants" binding:"omitempty,gt=0"`
	Status          *string `json:"status"`
}

type SquadOrganizer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Position string `json:"position"`
	Phone    string `json:"phone"`
}

type StudentSquadResponse struct {
	ID              string         `json:"id"`
	Organizer       SquadOrganizer `json:"organizer"`
	Title           string         `json:"title"`
	Description     *string        `json:"description,omitempty"`
	Profile         *string        `json:"profile,omitempty"`
	MaxParticipants int            `json:"max_participants"`
	CurrentCount    int            `json:"current_count"`
	Status          string         `json:"status"`
	StatusID        int            `json:"status_id"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}

type StudentSquadDetailResponse struct {
	StudentSquadResponse
	ParticipantIDs []string `json:"participant_ids"`
}

type SquadStatusResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type JoinSquadRequest struct {
	SquadID string `json:"squad_id" binding:"required"`
}

type AllSquadsResponse struct {
	Squads []StudentSquadResponse `json:"squads"`
	Total  int                    `json:"total"`
	Limit  int                    `json:"limit"`
	Offset int                    `json:"offset"`
}

type ApproveSquadRequest struct {
	Approved bool `json:"approved"`
}
