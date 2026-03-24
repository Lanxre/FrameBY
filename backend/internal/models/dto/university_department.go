package dto

type UniversityDepartmentResponse struct {
	ID             string  `json:"id"`
	UniversityName string  `json:"university_name"`
	DepartmentName string  `json:"department_name"`
	Address        *string `json:"address,omitempty"`
}
