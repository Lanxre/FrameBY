package dto

type UpdateUserProfileRequest struct {
	Login     string `json:"login"`
	Password string `json:"password"`
}

