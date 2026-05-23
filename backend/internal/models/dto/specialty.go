package dto

type SpecialtyResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID string `json:"department_id"`
}

type SpecialtiesListResponse struct {
	Specialties []SpecialtyResponse `json:"specialties"`
}
