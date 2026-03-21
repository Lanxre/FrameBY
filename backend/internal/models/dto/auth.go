package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  LoginUserDTO `json:"user"`
}

type LoginUserDTO struct {
	ID        uuid.UUID  `json:"id"`
	Email     string 	 `json:"email"`
	Login     string 	 `json:"login"`
	Role      string 	 `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UserDto struct {
	ID        uuid.UUID  `json:"id"`
	Email     string 	 `json:"email"`
	Login     string 	 `json:"login"`
	Role      string 	 `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	
	
	FullName     *string `json:"full_name"`
	Subrole      *string `json:"subrole"`    // только для BRSM
	EnterpriseID *uuid.UUID `json:"enterprise_id"` // только для Customer
}
