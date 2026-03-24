package dto

type EnterpriseResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address,omitempty"`
}
