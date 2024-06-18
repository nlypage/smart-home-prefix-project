package dto

// CreateConditioner is a struct that represents a DTO to create a new Conditioner.
type CreateConditioner struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

// UpdateConditioner is a struct that represents a DTO to update an existing Conditioner.
type UpdateConditioner struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Activated   bool   `json:"activated"`
	Temperature int    `json:"temperature"`
}
