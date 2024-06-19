package dto

// CreateAlarmSystem is a struct that represents a DTO to create a new AlarmSystem.
type CreateAlarmSystem struct {
	AccessToken string `json:"access_token" validate:"required,accessToken"`
	Name        string `json:"name" validate:"required,name"`
}

// UpdateAlarmSystem is a struct that represents a DTO to update an existing AlarmSystem.
type UpdateAlarmSystem struct {
	UUID        string `json:"uuid" validate:"required,uuid4"`
	AccessToken string `json:"access_token" validate:"required,accessToken"`
	Name        string `json:"name" validate:"required,name"`
	Activated   bool   `json:"activated" validate:"required"`
	Triggered   bool   `json:"triggered" validate:"required"`
}
