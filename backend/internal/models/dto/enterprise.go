package dto

type CreateEnterpriseRequest struct {
	Name    string  `json:"name" binding:"required"`
	Address *string `json:"address"`
}

type EnterpriseResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address,omitempty"`
}
