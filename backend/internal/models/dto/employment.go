package dto

type CreateEmploymentRequest struct {
	Title           string  `json:"title" binding:"required"`
	Description     *string `json:"description"`
	Requirements    *string `json:"requirements"`
	Salary          *string `json:"salary"`
	Schedule        *string `json:"schedule"`
	MaxParticipants int     `json:"max_participants" binding:"required,gt=0"`
}

type UpdateEmploymentRequest struct {
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	Requirements    *string `json:"requirements"`
	Salary          *string `json:"salary"`
	Schedule        *string `json:"schedule"`
	MaxParticipants *int    `json:"max_participants" binding:"omitempty,gt=0"`
	Status          *string `json:"status"`
}

type EmploymentRequestResponse struct {
	ID                     string  `json:"id"`
	EnterpriseID           string  `json:"enterprise_id"`
	EnterpriseName         string  `json:"enterprise_name"`
	EnterpriseAddress      *string `json:"enterprise_address,omitempty"`
	UniversityDepartmentID string  `json:"university_department_id"`
	UniversityName         string  `json:"university_name"`
	DepartmentName         string  `json:"department_name"`
	UniversityAddress      *string `json:"university_address,omitempty"`
	Title                  string  `json:"title"`
	Description            *string `json:"description,omitempty"`
	Requirements           *string `json:"requirements,omitempty"`
	Salary                 *string `json:"salary,omitempty"`
	Schedule               *string `json:"schedule,omitempty"`
	MaxParticipants        int     `json:"max_participants"`
	CurrentParticipants    int     `json:"current_participants"`
	Status                 string  `json:"status"`
	StatusID               int     `json:"status_id"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
}

type EmploymentRequestDetailResponse struct {
	EmploymentRequestResponse
	ParticipantUserIDs []string `json:"participant_user_ids"`
}

type EmploymentApplicationResponse struct {
	EmploymentRequestResponse
	ParticipantStatus    string `json:"participant_status"`
	ParticipantStatusID  int    `json:"participant_status_id"`
	ParticipantAppliedAt string `json:"participant_applied_at"`
}

type EmploymentParticipantResponse struct {
	ID           string  `json:"id"`
	RequestID    string  `json:"request_id"`
	UserID       string  `json:"user_id"`
	StudentName  string  `json:"student_name"`
	Status       string  `json:"status"`
	StatusID     int     `json:"status_id"`
	AppliedAt    string  `json:"applied_at"`
	ContractedAt *string `json:"contracted_at,omitempty"`
}

type EmploymentStatusResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmploymentParticipantStatusResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ApplyToEmploymentRequest struct {
	RequestID string `json:"request_id" binding:"required"`
}

type UpdateParticipantStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type AllEmploymentRequestsResponse struct {
	Requests []EmploymentRequestResponse `json:"requests"`
	Total    int                         `json:"total"`
	Limit    int                         `json:"limit"`
	Offset   int                         `json:"offset"`
}

type EmploymentParticipantsResponse struct {
	Participants []EmploymentParticipantResponse `json:"participants"`
	Total        int                             `json:"total"`
}

type ApproveEmploymentRequest struct {
	Approved bool `json:"approved"`
}
