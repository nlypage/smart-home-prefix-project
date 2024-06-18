package dto

// CreateHeatingControl is a struct that represents a DTO to create a new HeatingControl.
type CreateHeatingControl struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

// UpdateHeatingControl is a struct that represents a DTO to update an existing HeatingControl.
type UpdateHeatingControl struct {
	UUID        string `json:"uuid"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Activated   bool   `json:"activated"`
	Temperature int    `json:"temperature"`
}
