package dto

type CreateSpecialtyRequest struct {
	Name         string `json:"name" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
}

type SpecialtyResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID string `json:"department_id"`
}

type SpecialtiesListResponse struct {
	Specialties []SpecialtyResponse `json:"specialties"`
}
