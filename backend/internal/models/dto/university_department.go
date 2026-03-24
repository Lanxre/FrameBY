package dto

type CreateUniversityDepartmentRequest struct {
	UniversityName string  `json:"university_name" binding:"required"`
	DepartmentName string  `json:"department_name" binding:"required"`
	Address        *string `json:"address"`
}

type UniversityDepartmentResponse struct {
	ID             string  `json:"id"`
	UniversityName string  `json:"university_name"`
	DepartmentName string  `json:"department_name"`
	Address        *string `json:"address,omitempty"`
}
