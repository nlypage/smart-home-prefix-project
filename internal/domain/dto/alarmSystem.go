package dto

// CreateAlarmSystem is a struct that represents a DTO to create a new AlarmSystem.
type CreateAlarmSystem struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name" json:"name"`
}

// UpdateAlarmSystem is a struct that represents a DTO to update an existing AlarmSystem.
type UpdateAlarmSystem struct {
	ID          uint   `json:"ID"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Activated   bool   `json:"activated"`
	Triggered   bool   `json:"triggered"`
}
